package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Route("/ping", func(r chi.Router) {
			r.Get("/", s.handlePing)
		})

		r.Route("/account", func(r chi.Router) {
			r.Route("/{accountId}", func(r chi.Router) {
				r.Route("/workout", func(r chi.Router) {
					r.Get("/", s.handleGetWorkoutByDate)

					r.Route("/{workoutId}", func(r chi.Router) {
						r.Route("/session", func(r chi.Router) {
							r.Post("/", s.handlePostWorkoutSession)
						})
					})
				})
			})
		})

		r.Route("/session", func(r chi.Router) {
			r.Route("/{sessionId}", func(r chi.Router) {
				r.Patch("/", s.handlePatchSession)

				r.Route("/exercise", func(r chi.Router) {
					r.Route("/{exerciseId}", func(r chi.Router) {
						r.Post("/", s.handlePostExerciseLog)
					})
				})
			})
		})
	})
}
