package repository

import (
	"context"
	"database/sql"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/google/uuid"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (s *SessionRepo) CreateWorkoutSession(
	ctx context.Context,
	accountId uuid.UUID,
	workoutId uuid.UUID,
	session *model.Session) error {
	query := `INSERT INTO workout_session (account, plan_workouts, started_at)
				VALUES ($1, $2, $3)
				RETURNING id`

	return s.db.QueryRowContext(
		ctx,
		query,
		accountId,
		workoutId,
		session.StartedAt).Scan(&session.Id)
}

func (s *SessionRepo) UpdateSession(ctx context.Context, session *model.Session) error {
	query := `UPDATE workout_session
			SET completed_at = $1, 
    			total_duration = EXTRACT(EPOCH FROM ($2 - started_at)),
				total_calories_burned = $3,
    			notes = $4
			WHERE id = $5`

	_, err := s.db.ExecContext(ctx,
		query,
		session.CompletedAt,
		session.CompletedAt,
		session.TotalCaloriesBurned,
		session.Notes,
		session.Id)

	return err
}
