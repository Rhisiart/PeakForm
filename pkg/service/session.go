package service

import (
	"context"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/Rhisiart/PeakForm/pkg/repository"
)

type SessionService struct {
	SessionRepo repository.SessionRepository
}

func NewSessionService(sessionRepo repository.SessionRepository) *SessionService {
	return &SessionService{
		SessionRepo: sessionRepo,
	}
}

func (s *SessionService) UpdateSession(ctx context.Context, session *model.Session) error {
	return s.SessionRepo.UpdateSession(ctx, session)
}
