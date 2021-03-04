package user

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

// func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
// 	return json.NewEncoder(w).Encode(response)
// }

// func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
// 	var req CreateUserRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }

// func decodeNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
// 	var req GetUserRequest
// 	vars := mux.Vars(r)

// 	req = GetUserRequest{
// 		ID: vars["id"],
// 	}
// 	return req, nil
// }
