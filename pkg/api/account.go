package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (s *Server) handleGetWorkoutByDate(w http.ResponseWriter, r *http.Request) {
	accountId, err := uuid.Parse(chi.URLParam(r, "id"))
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

	workout, err := s.service.AccountService.GetWorkoutByDate(r.Context(), accountId.String(), date)

	if err != nil {
		slog.Error("Cannot query the db", "Error", err.Error())
		render.Render(w, r, NewError(
			err,
			"Unable to retrieve the workout for the day",
			http.StatusBadRequest,
		))
		return
	}

	render.Render(w, r, workout)
}
