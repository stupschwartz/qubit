package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/bmizerany/pq"
	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc/status"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/applications/lib/pgutils"
	"github.com/stupschwartz/qubit/core/computation"
	"github.com/stupschwartz/qubit/core/computation_status"
	computations_pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
)

const pubSubCoordinatorSubscriptionID = "coordinator"

var heartbeatInterval = 10 * time.Second

// TODO: What's a reasonable ack deadline? Probably the amount of time in which
// TODO: any calculation would finish, because we ack after it completes.
var messageAckInterval = 60 * time.Second

type Coordinator struct {
	PostgresClient *sqlx.DB
}

func validateLastComputationStatus(tx *sqlx.Tx, computationId string, lastComputationStatusId string) error {
	var compStatuses []computation_status.ComputationStatus
	err := pgutils.Select(&pgutils.SelectConfig{
		Args: []interface{}{computationId},
		// Strictly limit columns for performance
		Columns: []string{
			fmt.Sprintf("%v.id", computation_status.TableName),
		},
		ForClause: "FOR UPDATE",
		// Join to computation to ensure existence of computation and to lock it, as well
		JoinClause: fmt.Sprintf("INNER JOIN %v ON %v.computation_id = %v.id",
			computation.TableName, computation_status.TableName, computation.TableName),
		Limit:         1,
		OrderByClause: fmt.Sprintf("ORDER BY %v.created_at DESC", computation_status.TableName),
		Out:           &compStatuses,
		Table:         computation_status.TableName,
		Tx:            tx,
		WhereClause:   "WHERE computation_id=$1",
	})
	if err != nil {
		return err
	} else if len(compStatuses) == 0 {
		return errors.Wrapf(err, "No computation statuses exist for computation with ID: %v", computationId)
	} else if compStatuses[0].Id != lastComputationStatusId {
		return errors.Wrapf(err, "Computation status %v does not match newest status %v for computation %v",
			lastComputationStatusId, compStatuses[0].Id, computationId)
	}
	return nil
}

func (c *Coordinator) heartbeat(tkr *time.Ticker, firstCompStatus *computation_status.ComputationStatus) {
	computationId := firstCompStatus.ComputationId
	// Keep a reference to the last computation status created from heartbeat
	lastCompStatus := firstCompStatus
	// First tick occurs after duration
	for range tkr.C {
		tx, err := c.PostgresClient.Beginx()
		if err != nil {
			log.Println(errors.Wrap(err, "Failed to begin transaction"))
			continue
		}
		err = validateLastComputationStatus(
			tx,
			lastCompStatus.ComputationId,
			lastCompStatus.Id,
		)
		if err != nil {
			log.Println(err)
			return
		}
		newCompStatus := computation_status.New(
			computationId,
			computation_status.ComputationStatusInProgress,
		)
		err = apiutils.Create(&apiutils.CreateConfig{
			Object: &newCompStatus,
			Table:  computation_status.TableName,
			Tx:     tx,
		})
		if err != nil {
			log.Println(errors.Wrapf(err, "Could not create new computation status: %v", newCompStatus))
			continue
		}
		err = tx.Commit()
		if err != nil {
			log.Println(errors.Wrap(err, "Failed to commit transaction"))
			continue
		}
		lastCompStatus = newCompStatus
	}
}

func (c *Coordinator) subscriptionHandler(ctx context.Context, msg *pubsub.Message) {
	var msgPBCompStatus computations_pb.ComputationStatus
	err := proto.Unmarshal(msg.Data, &msgPBCompStatus)
	if err != nil {
		log.Println(err)
		msg.Ack()
		return
	}
	msgCompStatus := computation_status.NewFromProto(msgPBCompStatus)
	// Only proceed for created or requeued computation statuses
	if msgCompStatus.Status != computation_status.ComputationStatusCreated &&
		msgCompStatus.Status != computation_status.ComputationStatusRequeued {
		log.Printf("Computation %v cannot be started from status %v", msgCompStatus.ComputationId, msgCompStatus.Status)
		msg.Ack()
		return
	}
	tx, err := c.PostgresClient.Beginx()
	if err != nil {
		log.Println(err)
		msg.Nack()
		return
	}
	err = validateLastComputationStatus(
		tx,
		msgCompStatus.ComputationId,
		msgCompStatus.Id,
	)
	if err != nil {
		log.Println(err)
		msg.Ack()
		return
	}
	newCompStatus := computation_status.New(
		msgCompStatus.ComputationId,
		computation_status.ComputationStatusStarted,
	)
	err = apiutils.Create(&apiutils.CreateConfig{
		Object: &newCompStatus,
		Table:  computation_status.TableName,
		Tx:     tx,
	})
	if err != nil {
		log.Println(err)
		msg.Nack()
		return
	}
	// Set up heartbeat before committing transaction to ensure computation won't get lost
	heartbeatTicker := time.NewTicker(heartbeatInterval)
	go c.heartbeat(heartbeatTicker, &newCompStatus)
	err = tx.Commit()
	if err != nil {
		heartbeatTicker.Stop()
		log.Println(err)
		msg.Nack()
		return
	}
	log.Println("Processing computation:", msgCompStatus.ComputationId)
	///////////////////////
	// TODO: Business logic
	///////////////////////
	log.Println("Acknowledging message:", msg)
	msg.Ack()
	log.Println("Stopping heartbeat for computation:", msgCompStatus.ComputationId)
	heartbeatTicker.Stop()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}
	projID := os.Getenv("GOOGLE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_PROJECT_ID"`)
	}
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal(`You need to set the environment variable "POSTGRES_URL"`)
	}
	parsedURL, err := pq.ParseURL(postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	postgresClient, err := sqlx.Open("postgres", parsedURL)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	pubSubClient, err := pubsub.NewClient(ctx, projID, serviceCredentials)
	for err != nil {
		log.Printf("Could not create pubsub client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		pubSubClient, err = pubsub.NewClient(ctx, projID, serviceCredentials)
	}
	topic, err := pubSubClient.CreateTopic(ctx, computation_status.PubSubTopicID)
	if err != nil {
		// 409 ALREADY_EXISTS is an inevitable and harmless error
		// https://cloud.google.com/pubsub/docs/reference/error-codes
		if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != 409 {
			return nil, errors.Wrapf(err, "Failed to get-or-create topic %v", computation_status.PubSubTopicID)
		}
	}
	subscription, err := pubSubClient.CreateSubscription(ctx, pubSubCoordinatorSubscriptionID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: messageAckInterval,
	})
	if err != nil {
		// 409 ALREADY_EXISTS is an inevitable and harmless error
		// https://cloud.google.com/pubsub/docs/reference/error-codes
		if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != 409 {
			return nil, errors.Wrapf(err, "Failed to get-or-create subscription %v", pubSubCoordinatorSubscriptionID)
		}
	}
	coordinator := Coordinator{PostgresClient: postgresClient}
	err = subscription.Receive(ctx, coordinator.subscriptionHandler)
	if err != nil {
		log.Fatal(err)
	}
}
