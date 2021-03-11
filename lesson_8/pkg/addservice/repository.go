package addservice

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

// Repository interface
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}

// ErrRepo repository error
var ErrRepo = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo creates new repo
func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `
		INSERT INTO users (id, name)
		VALUES ($1, $2)`

	if user.Name == "" {
		return ErrRepo
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var name string
	err := repo.db.QueryRow("SELECT name FROM users WHERE id=$1", id).Scan(&name)
	if err != nil {
		return "", ErrRepo
	}

	return name, nil
}
