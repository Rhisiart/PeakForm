package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/google/uuid"
)

type AccountRepository interface {
	FindWorkoutByDate(ctx context.Context, accountId uuid.UUID, dayOfWeek int, date time.Time) (*model.Workout, error)
	CreateWorkoutSession(ctx context.Context, accountId uuid.UUID, workoutId uuid.UUID, session *model.Session) error
}
type SessionRepository interface {
	UpdateSession(ctx context.Context, session *model.Session) error
}
type Repository struct {
	AccountRepository AccountRepository
	SessionRepository SessionRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AccountRepository: NewAccountRepo(db),
		SessionRepository: NewSessionRepo(db),
	}
}
