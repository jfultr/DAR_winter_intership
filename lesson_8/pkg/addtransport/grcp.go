package addtransport

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pb"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pkg/addendpoint"
)

type gRPCServer struct {
	createUser gt.Handler
	getUser    gt.Handler
}

// NewGRPCServer creates new server
func NewGRPCServer(ctx context.Context, endpoints addendpoint.Endpoints) pb.UserServiceServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			endpoints.CreateUserEndpoint,
			decodeGRCPUserReq,
			encodeGRCPUserResponse,
		),
		getUser: gt.NewServer(
			endpoints.GetUserEndpoint,
			decodeGRCPNameReq,
			encodeGRCPNameResponse,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResponse), nil
}

func (s *gRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetUserResponse), nil
}

func decodeGRCPUserReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return addendpoint.CreateUserReq{Name: req.Name}, nil
}

func encodeGRCPUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.CreateUserResp)
	return &pb.CreateUserResponse{Ok: resp.Ok}, nil
}

func decodeGRCPNameReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return addendpoint.GetUserReq{ID: req.ID}, nil
}

func encodeGRCPNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.GetUserResp)
	return &pb.GetUserResponse{Name: resp.Name}, nil
}
