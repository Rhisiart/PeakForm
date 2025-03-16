package repository

import (
	"context"
	"database/sql"
	"log/slog"
)

type AccountRepo struct {
	db *sql.DB
}

func NewWorkout(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (wr *AccountRepo) FindWorkoutForDay(
	ctx context.Context,
	accountId string,
	week int,
	dayOfWeek int) {
	query := `SELECT we.exercise, we.reps, we.sets, we.weight, we.rest, we.notes
    			FROM workout_exercise AS we
    			JOIN workout AS w ON w.id = we.workout
    			JOIN plan_workouts AS pw ON pw.workout = w.id
    			JOIN training_plan AS pl ON pl.id = pw.training_plan
    			JOIN account_training_plan AS atp ON atp.training_plan = pl.id
    			WHERE atp.account = '$1' AND atp.status = 'Active' AND pw.week = $2 AND pw.day_of_week = $3
    			ORDER BY we.order_number ASC`

	rows, err := wr.db.QueryContext(ctx, query, accountId, week, dayOfWeek)

	if err != nil {
		slog.Error("Cannot get the workout for the day", "Error", err.Error())
		return err
	}

}
