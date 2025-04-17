package service

import (
	"context"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/Rhisiart/PeakForm/pkg/repository"
	"github.com/google/uuid"
)

type IExerciseLogService interface {
	CreateExerciseLog(ctx context.Context, sessionId uuid.UUID, exerciseId uuid.UUID, exerciseLog *model.ExerciseLog) error
}
type IAccountService interface {
	GetWorkoutByDate(ctx context.Context, accountId uuid.UUID, date string) (*model.Workout, error)
}
type ISessionService interface {
	CreateWorkoutSession(ctx context.Context, accoutId uuid.UUID, workoutId uuid.UUID, session *model.Session) error
	UpdateSession(ctx context.Context, session *model.Session) error
}
type Service struct {
	AccountService     IAccountService
	SessionService     ISessionService
	ExerciseLogService IExerciseLogService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AccountService:     NewAccontService(repo.AccountRepository),
		SessionService:     NewSessionService(repo.SessionRepository),
		ExerciseLogService: NewExerciseLogService(repo.ExerciseLogRepository),
	}
}
