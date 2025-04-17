package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/google/uuid"
)

type ExerciseLogRepository interface {
	CreateExerciseLog(ctx context.Context, sessionId uuid.UUID, exerciseId uuid.UUID, exerciseLog *model.ExerciseLog) error
}

type AccountRepository interface {
	FindWorkoutByDate(ctx context.Context, accountId uuid.UUID, dayOfWeek int, date time.Time) (*model.Workout, error)
}
type SessionRepository interface {
	CreateWorkoutSession(ctx context.Context, accountId uuid.UUID, workoutId uuid.UUID, session *model.Session) error
	UpdateSession(ctx context.Context, session *model.Session) error
}
type Repository struct {
	AccountRepository     AccountRepository
	SessionRepository     SessionRepository
	ExerciseLogRepository ExerciseLogRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AccountRepository:     NewAccountRepo(db),
		SessionRepository:     NewSessionRepo(db),
		ExerciseLogRepository: NewExerciseLogRepo(db),
	}
}
