package model

import (
	"net/http"

	"github.com/google/uuid"
)

type Workout struct {
	Id               uuid.UUID   `json:"id"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	WorkoutType      string      `json:"workoutType"`
	Difficulty       string      `json:"difficulty"`
	CaloriesEstimate int         `json:"caloriesEstimate"`
	Exercises        []*Exercise `json:"exercises"`
}

func (wkt *Workout) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
