package user

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
)

type mustIm struct {
}
type gRPCServer struct {
	createUser gt.Handler
	getUser    gt.Handler
	mi         mustIm
}

// NewGRPCServer creates new server
func NewGRPCServer(ctx context.Context, endpoints Endpoints) UserServiceServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			endpoints.CreateUser,
			decodeUserReq,
			encodeUserResponse,
		),
		getUser: gt.NewServer(
			endpoints.GetUser,
			decodeNameReq,
			encodeNameResponse,
		),
		mi: mustIm{},
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateUserResponse), nil
}

func (s *gRPCServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetUserResponse), nil
}

func decodeUserReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*CreateUserRequest)
	return CreateUserReq{Name: req.Name}, nil
}

func encodeUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CreateUserResp)
	return &CreateUserResponse{Ok: resp.Ok}, nil
}

func decodeNameReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*GetUserRequest)
	return GetUserReq{ID: req.ID}, nil
}

func encodeNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetUserResp)
	return &GetUserResponse{Name: resp.Name}, nil
}
