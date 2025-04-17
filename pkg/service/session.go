package service

import (
	"context"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/Rhisiart/PeakForm/pkg/repository"
	"github.com/google/uuid"
)

type SessionService struct {
	SessionRepo repository.SessionRepository
}

func NewSessionService(sessionRepo repository.SessionRepository) *SessionService {
	return &SessionService{
		SessionRepo: sessionRepo,
	}
}

func (s *SessionService) CreateWorkoutSession(
	ctx context.Context,
	accoutId uuid.UUID,
	workoutId uuid.UUID,
	session *model.Session) error {
	return s.SessionRepo.CreateWorkoutSession(ctx, accoutId, workoutId, session)
}

func (s *SessionService) UpdateSession(ctx context.Context, session *model.Session) error {
	return s.SessionRepo.UpdateSession(ctx, session)
}
