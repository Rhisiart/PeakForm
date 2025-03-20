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

func (s *Server) handleGetWorkoutByDate(w http.ResponseWriter, r *http.Request) {
	accountId, err := uuid.Parse(chi.URLParam(r, "accountId"))
	date := r.URL.Query().Get("date")

	slog.Warn("Starting to get the workout for day", "AccountId", accountId.String(), "Date", date)

	if err != nil || date == "" {
		slog.Error("Invalid query parameters", "AccountId", accountId.String(), "Date", date)
		render.Render(w, r, NewError(
			errors.New("invalid query parameters"),
			"Invalid request, the accountId or time are invalid",
			http.StatusBadRequest,
		))
		return
	}

	workout, err := s.service.AccountService.GetWorkoutByDate(r.Context(), accountId, date)

	if err != nil {
		slog.Error("Cannot query the db", "Error", err.Error())
		render.Render(w, r, NewError(
			err,
			"Unable to retrieve the workout for the day",
			http.StatusInternalServerError,
		))
		return
	}

	render.Render(w, r, workout)
}

func (s *Server) handlePostWorkoutSession(w http.ResponseWriter, r *http.Request) {
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
	err = s.service.AccountService.CreateWorkoutSession(r.Context(), accountId, workoutId, session)

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
