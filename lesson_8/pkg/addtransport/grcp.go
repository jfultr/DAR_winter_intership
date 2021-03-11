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
			decodeUserReq,
			encodeUserResponse,
		),
		getUser: gt.NewServer(
			endpoints.GetUserEndpoint,
			decodeNameReq,
			encodeNameResponse,
		),
	}
}

// // NewGRPCClient returns client
// func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) addservice.Service {

// 	// Each individual endpoint is an grpc/transport.Client (which implements
// 	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
// 	// made your own client library, you'd do this work there, so your server
// 	// could rely on a consistent set of client behavior.
// 	var createUserEndpoint endpoint.Endpoint
// 	{
// 		createUserEndpoint = gt.NewClient(
// 			conn,
// 			"pb.UserService",
// 			"CreateUser",
// 			decodeUserReq,
// 			encodeUserResponse,
// 			pb.CreateUserResponse{},
// 		).Endpoint()

// 	}

// 	// The Concat endpoint is the same thing, with slightly different
// 	// middlewares to demonstrate how to specialize per-endpoint.
// 	var getUserEndpoint endpoint.Endpoint
// 	{
// 		getUserEndpoint = gt.NewClient(
// 			conn,
// 			"pb.UserService",
// 			"GetUser",
// 			decodeNameReq,
// 			encodeNameResponse,
// 			pb.GetUserResponse{},
// 		).Endpoint()
// 	}

// 	// Returning the endpoint.Set as a service.Service relies on the
// 	// endpoint.Set implementing the Service methods. That's just a simple bit
// 	// of glue code.
// 	return addendpoint.Endpoints{
// 		CreateUserEndpoint: createUserEndpoint,
// 		GetUserEndpoint:    getUserEndpoint,
// 	}
// }

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

func decodeUserReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return addendpoint.CreateUserReq{Name: req.Name}, nil
}

func encodeUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.CreateUserResp)
	return &pb.CreateUserResponse{Ok: resp.Ok}, nil
}

func decodeNameReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return addendpoint.GetUserReq{ID: req.ID}, nil
}

func encodeNameResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.GetUserResp)
	return &pb.GetUserResponse{Name: resp.Name}, nil
}
