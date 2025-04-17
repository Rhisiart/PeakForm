package model

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type ExerciseLog struct {
	Id uuid.UUID `json:"id"`
	//SessionId     uuid.UUID `json:"sessionId"`
	//ExerciseId    uuid.UUID `json:"exerciseId"`
	Sets          int    `json:"sets"`
	RepsDone      int    `json:"repsDone"`
	WeightUsed    int    `json:"weightUsed"`
	Feeling       string `json:"feeling"`
	RepsInReserve int    `json:"repsInReserve"`
	Completed     bool   `json:"completed"`
	Notes         string `json:"notes"`
}

func (e *ExerciseLog) Bind(r *http.Request) error {
	if e.Feeling != "Easy" && e.Feeling != "Moderate" && e.Feeling != "Hard" && e.Feeling != "Failed" {
		return errors.New("invalid feeling value")
	}

	if e.Sets < 0 || e.RepsDone < 0 || e.WeightUsed < 0 || e.RepsInReserve < 0 {
		return errors.New("invalid numeric value, should be greater than zero")
	}

	return nil
}

func (e *ExerciseLog) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
