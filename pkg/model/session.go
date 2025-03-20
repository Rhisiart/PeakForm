package model

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id                  uuid.UUID `json:"id"`
	StartedAt           time.Time `json:"startedAt"`
	CompletedAt         time.Time `json:"completedAt"`
	TotalDuration       int       `json:"totalDuration,omitempty"`
	TotalCaloriesBurned int       `json:"totalCaloriesBurned,omitempty"`
	Notes               string    `json:"notes,omitempty"`
}

func (s *Session) Bind(r *http.Request) error {
	//Need to find a way to validate multiples methods, because same properties are optiminal
	/*if s.StartedAt.IsZero() {
		return errors.New("the started time is required")
	}

	if s.CompletedAt.IsZero() {
		return errors.New("invalid completedAt value")
	}
	if s.TotalCaloriesBurned < 0 {
		return errors.New("totalCaloriesBurned must be positive")
	}*/

	return nil
}

func (s *Session) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
