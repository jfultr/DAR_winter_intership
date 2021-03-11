package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/examples/addsvc/pb"
	"github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/examples/addsvc/pkg/addservice"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
)

type gRPCServer struct {
	createUser gt.Handler
	getUser    gt.Handler
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
	}
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) addservice.Service {

	// Each individual endpoint is an grpc/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
	// made your own client library, you'd do this work there, so your server
	// could rely on a consistent set of client behavior.
	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = gt.NewClient(
			conn,
			"pb.Add",
			"CreateUser",
			decodeUserReq,
			encodeUserResponse,
			pb.SumReply{},
		).Endpoint()

	}

	// The Concat endpoint is the same thing, with slightly different
	// middlewares to demonstrate how to specialize per-endpoint.
	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = gt.NewClient(
			conn,
			"pb.Add",
			"GetUser",
			decodeNameReq,
			encodeNameResponse,
			pb.ConcatReply{},
		).Endpoint()
	}

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return addendpoint.Set{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
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
