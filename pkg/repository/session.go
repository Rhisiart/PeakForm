package repository

import (
	"context"
	"database/sql"

	"github.com/Rhisiart/PeakForm/pkg/model"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
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
