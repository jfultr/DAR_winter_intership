package addservice

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// User interface
type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Service interface
type Service interface {
	CreateUser(ctx context.Context, name string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

// NewService creates new service
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

type service struct {
	repostory Repository
	logger    log.Logger
}

func (s service) CreateUser(ctx context.Context, name string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:   id,
		Name: name,
	}

	if err := s.repostory.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	name, err := s.repostory.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return name, nil
}
