package api

import (
	"log/slog"
	"net/http"

	"github.com/Rhisiart/PeakForm/pkg/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (s *Server) handlePatchSession(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
