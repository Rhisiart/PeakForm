package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/api/v1.0", func(r chi.Router) {
		r.Route("/ping", func(r chi.Router) {
			r.Get("/", s.handlePing)
		})

		r.Route("/account", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Route("/workout", func(r chi.Router) {
					r.Get("/", s.handleGetWorkoutByDate)
				})
			})
		})
	})
}
