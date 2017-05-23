package scenes

import (
	"github.com/stupschwartz/qubit/server/env"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/scene"
	scenes_pb "github.com/stupschwartz/qubit/server/protos/scenes"
	"github.com/golang/protobuf/ptypes/empty"
	"math/rand"
	"time"
	"google.golang.org/grpc"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	env *env.Env
}

func (s *Server) List(ctx context.Context, in *empty.Empty) (*scenes_pb.ScenesList, error) {
	var scenes scene.Scenes
	_, err := s.env.DatastoreClient.GetAll(ctx, datastore.NewQuery("Scene"), &scenes)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	return scenes.ToProto(), nil
}

func (s *Server) Get(ctx context.Context, in *scenes_pb.GetSceneRequest) (*scenes_pb.Scene, error) {
	sceneKey := datastore.IDKey("Scene", in.SceneId, nil)

	var existingScene scene.Scene
	if err := s.env.DatastoreClient.Get(ctx, sceneKey, &existingScene); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingScene.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *scenes_pb.CreateSceneRequest) (*scenes_pb.Scene, error) {
	sceneId := r.Int63()
	sceneKey := datastore.IDKey("Scene", sceneId, nil)

	newScene := scene.NewSceneFromProto(in.Scene)
	newScene.Id = sceneId

	if _, err := s.env.DatastoreClient.Put(ctx, sceneKey, newScene); err != nil {
		return nil, errors.Wrapf(err, "Failed to put node %v", newScene.Id)
	}

	return newScene.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *scenes_pb.UpdateSceneRequest) (*scenes_pb.Scene, error) {
	sceneKey := datastore.IDKey("Scene", in.SceneId, nil)

	newScene := scene.NewSceneFromProto(in.Scene)

	_, err := s.env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingScene scene.Scene
		if err := tx.Get(sceneKey, &existingScene); err != nil {
			return errors.Wrapf(err, "Failed to get scene in tx %v", existingScene)
		}

		existingScene.Name = newScene.Name

		_, err := tx.Put(sceneKey, &existingScene)
		if err != nil {
			return errors.Wrapf(err, "Failed to put scene in tx %v", existingScene)
		}

		newScene = existingScene

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update scene")
	}

	return newScene.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *scenes_pb.DeleteSceneRequest) (*empty.Empty, error) {
	sceneKey := datastore.IDKey("Scene", in.SceneId, nil)

	if err := s.env.DatastoreClient.Delete(ctx, sceneKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted scene by id: %v", in.SceneId)
	}

	return &empty.Empty{}, nil
}

func newServer(e *env.Env) *Server {
	return &Server{
		env: e,
	}
}

func Register(server *grpc.Server, e *env.Env) {
	scenes_pb.RegisterScenesServer(server, newServer(e))
}
