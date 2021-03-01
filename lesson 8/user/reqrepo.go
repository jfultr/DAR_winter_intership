package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// CreateUserRequest struct
	CreateUserRequest struct {
		Name string `json:"name"`
	}

	// CreateUserResponse struct
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	// GetUserRequest struct
	GetUserRequest struct {
		ID string `json:"id"`
	}

	// GetUserResponse struct
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		ID: vars["id"],
	}
	return req, nil
}
