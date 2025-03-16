package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/workout", func(r chi.Router) {
		s.router.Get("/{week}", s.handleGetWorkoutForWeek)
	})
}
