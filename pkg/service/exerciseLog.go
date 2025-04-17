package service

import (
	"context"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/Rhisiart/PeakForm/pkg/repository"
	"github.com/google/uuid"
)

type ExerciseLogService struct {
	exerciseLogRepo repository.ExerciseLogRepository
}

func NewExerciseLogService(exerciseLogRepo repository.ExerciseLogRepository) *ExerciseLogService {
	return &ExerciseLogService{
		exerciseLogRepo: exerciseLogRepo,
	}
}

func (e *ExerciseLogService) CreateExerciseLog(
	ctx context.Context,
	sessionId uuid.UUID,
	exerciseId uuid.UUID,
	exerciseLog *model.ExerciseLog,
) error {
	return e.exerciseLogRepo.CreateExerciseLog(ctx, sessionId, exerciseId, exerciseLog)
}
