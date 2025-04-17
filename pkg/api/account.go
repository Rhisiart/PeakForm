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
	slog.Warn("Received a request to get the workout by date")
	accountId, err := uuid.Parse(chi.URLParam(r, "accountId"))
	date := r.URL.Query().Get("date")

	if err != nil || date == "" {
		slog.Error("Invalid query parameters", "AccountId", accountId.String(), "Date", date)
		render.Render(w, r, NewError(
			errors.New("invalid query parameters"),
			"Invalid request, the accountId or time are invalid",
			http.StatusBadRequest,
		))
		return
	}

	slog.Warn("Starting to get the workout for day", "AccountId", accountId.String(), "Date", date)
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

	slog.Warn("Successfull to retrieve the workout for the date", "AccountId", accountId, "Date", date)
	render.Render(w, r, workout)
}
