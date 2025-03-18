package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rhisiart/PeakForm/pkg/model"
)

type AccountRepository interface {
	FindWorkoutByDate(ctx context.Context, accountId string, dayOfWeek int, date time.Time) (*model.Workout, error)
}
type Repository struct {
	AccountRepository AccountRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AccountRepository: NewAccountRepo(db),
	}
}
