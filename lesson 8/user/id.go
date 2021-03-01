package user

import "context"

// User interface
type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Repository interface
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}
