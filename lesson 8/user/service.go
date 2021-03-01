package user

import "context"

// Service interface
type Service interface {
	CreateUser(ctx context.Context, name string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
