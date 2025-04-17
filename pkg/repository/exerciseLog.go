package repository

import (
	"context"
	"database/sql"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/google/uuid"
)

type ExerciseLogRepo struct {
	db *sql.DB
}

func NewExerciseLogRepo(db *sql.DB) *ExerciseLogRepo {
	return &ExerciseLogRepo{
		db: db,
	}
}

func (e *ExerciseLogRepo) CreateExerciseLog(
	ctx context.Context,
	sessionId uuid.UUID,
	exerciseId uuid.UUID,
	exerciseLog *model.ExerciseLog,
) error {
	query := `INSERT INTO exercise_log (workout_session, exercise, sets, reps_done, weight_used, feeling, reps_in_reserve, completed, notes)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id`

	return e.db.QueryRowContext(
		ctx,
		query,
		sessionId,
		exerciseId,
		exerciseLog.Sets,
		exerciseLog.RepsDone,
		exerciseLog.WeightUsed,
		exerciseLog.Feeling,
		exerciseLog.RepsInReserve,
		exerciseLog.Completed,
		exerciseLog.Notes).Scan(&exerciseLog.Id)
}
