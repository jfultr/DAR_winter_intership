package addendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pkg/addservice"
)

// Endpoints struct
type Endpoints struct {
	CreateUserEndpoint endpoint.Endpoint
	GetUserEndpoint    endpoint.Endpoint
}

// MakeEndpoints makes enpoint
func MakeEndpoints(s addservice.Service) Endpoints {
	return Endpoints{
		CreateUserEndpoint: makeCreateUserEndpoint(s),
		GetUserEndpoint:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s addservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserReq)
		ok, err := s.CreateUser(ctx, req.Name)
		return CreateUserResp{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s addservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserReq)
		name, err := s.GetUser(ctx, req.ID)

		return GetUserResp{
			Name: name,
		}, err
	}
}

type (
	// CreateUserReq struct
	CreateUserReq struct {
		Name string
	}

	// CreateUserResp struct
	CreateUserResp struct {
		Ok string
	}

	// GetUserReq struct
	GetUserReq struct {
		ID string
	}

	// GetUserResp struct
	GetUserResp struct {
		Name string
	}
)
