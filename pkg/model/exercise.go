package model

type Exercise struct {
	Id     string `json:"id"`
	Reps   int    `json:"reps"`
	Sets   int    `json:"sets"`
	Weight int    `json:"weight"`
	Rest   int    `json:"rest"`
	Notes  string `json:"notes"`
}
