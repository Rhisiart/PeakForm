package model

import "github.com/google/uuid"

type Exercise struct {
	Id     uuid.UUID `json:"id"`
	Reps   int       `json:"reps"`
	Sets   int       `json:"sets"`
	Weight int       `json:"weight"`
	Rest   int       `json:"rest"`
	Notes  string    `json:"notes"`
}
