package service

import (
	"context"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/Rhisiart/PeakForm/pkg/repository"
)

type IAccountService interface {
	GetWorkoutByDate(ctx context.Context, accountId string, date string) (*model.Workout, error)
}
type Service struct {
	AccountService IAccountService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AccountService: NewAccontService(repo.AccountRepository),
	}
}
