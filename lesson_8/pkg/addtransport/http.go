package addtransport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jfultr/DAR_winter_intership/lesson_8/pkg/addendpoint"
)

// NewHTTPServer creates new server
func NewHTTPServer(ctx context.Context, endpoints addendpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		AuthMiddleware("jwt", endpoints.CreateUserEndpoint),
		decodeHTTPUserReq,
		encodeHTTPResponse,
		httptransport.ServerBefore(GetContext("jwt")),
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUserEndpoint,
		decodeHTTPNameReq,
		encodeHTTPResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addendpoint.CreateUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addendpoint.GetUserReq
	vars := mux.Vars(r)

	req = addendpoint.GetUserReq{
		ID: vars["id"],
	}
	return req, nil
}
