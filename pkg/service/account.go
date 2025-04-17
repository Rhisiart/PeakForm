package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/Rhisiart/PeakForm/pkg/repository"
	"github.com/google/uuid"
)

type AccountService struct {
	AccountRepo repository.AccountRepository
}

func NewAccontService(accountRepo repository.AccountRepository) *AccountService {
	return &AccountService{
		AccountRepo: accountRepo,
	}
}

func (a *AccountService) GetWorkoutByDate(
	ctx context.Context,
	accountId uuid.UUID,
	date string) (*model.Workout, error) {
	dt, err := time.Parse(time.DateOnly, date)
	weekDay := int(dt.Weekday())

	if err != nil {
		slog.Error("Unable to parse the date", "Error", err)
		return nil, err
	}

	slog.Warn("Querying the workouts", "AccountId", accountId, "Date", date, "WeekDay", weekDay)

	return a.AccountRepo.FindWorkoutByDate(ctx, accountId, weekDay, dt)
}
