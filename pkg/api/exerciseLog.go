package api

import (
	"log/slog"
	"net/http"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (s *Server) handlePostExerciseLog(w http.ResponseWriter, r *http.Request) {
	slog.Warn("Received a request to Create Exercise Log")
	sessionId, err := uuid.Parse(chi.URLParam(r, "sessionId"))
	exerciseId, errExerciseId := uuid.Parse(chi.URLParam(r, "exerciseId"))

	if err != nil || errExerciseId != nil {
		slog.Error("Cannot parse the ids", "SessionId", sessionId, "ExerciseId", exerciseId)
		render.Render(w, r, NewError(
			err,
			"Invalid session id or exercise id",
			http.StatusBadRequest,
		))
		return
	}

	exerciseLog := &model.ExerciseLog{}
	if err := render.Bind(r, exerciseLog); err != nil {
		slog.Error("Invalid request body", "Error", err)
		render.Render(w, r, NewError(
			err,
			"Invalid request body",
			http.StatusBadRequest,
		))
		return
	}

	slog.Warn("Query the db to create the exercise log")
	err = s.service.ExerciseLogService.CreateExerciseLog(r.Context(), sessionId, exerciseId, exerciseLog)

	if err != nil {
		slog.Error("Couldnt create the exercise log", "SessionId", sessionId, "ExerciseId", exerciseId)
		render.Render(w, r, NewError(
			err,
			"Couldnt create the exercise log",
			http.StatusInternalServerError,
		))
		return
	}

	slog.Warn("Successfully created the exercise log")
	render.Render(w, r, exerciseLog)
}
