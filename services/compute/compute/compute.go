package compute

import (
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/parameter"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Server struct {
}

func (s *Server) RenderImage(ctx context.Context, in *compute_pb.RenderImageRequest) (*compute_pb.RenderImageResponse, error) {
	op, err := operator.GetOperation(in.Operator.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Operation for rendering, %v", in.Operator.Type)
	}

	params := parameter.NewParametersFromProto(in.Operator.Parameters)

	imagePlane, err := op.Process(nil, params, in.BoundingBox.StartX, in.BoundingBox.StartY, in.BoundingBox.EndX, in.BoundingBox.EndY)


	return &compute_pb.RenderImageResponse{ImagePlane: imagePlane.ToProto()}, nil
}


func Register(grpcServer *grpc.Server) {
	compute_pb.RegisterComputeServer(grpcServer, &Server{})
}