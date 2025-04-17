package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (s *Server) handlePostWorkoutSession(w http.ResponseWriter, r *http.Request) {
	slog.Warn("Received a request to create a workout session")
	accountId, err := uuid.Parse(chi.URLParam(r, "accountId"))
	workoutId, errWId := uuid.Parse(chi.URLParam(r, "workoutId"))

	if err != nil || errWId != nil {
		slog.Error("Invalid Id's", "AccountId", accountId.String(), "WorkoutId", workoutId.String())
		render.Render(w, r, NewError(
			errors.New("invalid ids"),
			"Invalid request, the accountId or workoutId are incorrect",
			http.StatusBadRequest,
		))
		return
	}

	session := &model.Session{}
	if err = render.Bind(r, session); err != nil {
		slog.Error("Invalid request body", "Error", err.Error())
		render.Render(w, r, NewError(
			err,
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	slog.Warn("Creating a session for workout", "AccountId", accountId, "WorkoutId", workoutId, "Date", session.StartedAt)
	err = s.service.SessionService.CreateWorkoutSession(r.Context(), accountId, workoutId, session)

	if err != nil {
		slog.Error("Failed to create the session in the database", "Error", err.Error())
		render.Render(w, r, NewError(
			err,
			"Failed to create the session in the database",
			http.StatusInternalServerError,
		))
		return
	}

	slog.Warn("Successfully create the session", "Id", session.Id)
	render.Render(w, r, session)
}

func (s *Server) handlePatchSession(w http.ResponseWriter, r *http.Request) {
	slog.Warn("Received a request to Patch a session")
	sessionId, err := uuid.Parse(chi.URLParam(r, "sessionId"))

	if err != nil {
		slog.Error("The session id isnt a uuid", "SessionId", sessionId)
		render.Render(w, r, NewError(
			err,
			"Invalid uuid for the session Id",
			http.StatusBadRequest,
		))
		return
	}

	session := &model.Session{
		Id: sessionId,
	}

	if errBodyReq := render.Bind(r, session); errBodyReq != nil {
		slog.Error("Error in the body request", "Error", errBodyReq)
		render.Render(w, r, NewError(
			err,
			"Invalid body request",
			http.StatusBadRequest,
		))
		return
	}

	slog.Warn(
		"Querying the db to update the Session",
		"Id",
		session.Id,
		"CompletedAt",
		session.CompletedAt,
		"TotalCaloriesBurned",
		session.TotalCaloriesBurned,
		"Notes",
		session.Notes)
	err = s.service.SessionService.UpdateSession(r.Context(), session)

	if err != nil {
		slog.Error(
			"Unable to update the session",
			"Id",
			session.Id,
			"CompletedAt",
			session.CompletedAt,
			"TotalCaloriesBurned",
			session.TotalCaloriesBurned,
			"Notes",
			session.Notes,
		)
		render.Render(w, r, NewError(
			err,
			"Unable to update the Session",
			http.StatusInternalServerError,
		))
		return
	}

	slog.Warn("Successfully updated the Session")
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
