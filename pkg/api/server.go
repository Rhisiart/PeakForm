package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rhisiart/PeakForm/pkg/config"
	"github.com/Rhisiart/PeakForm/pkg/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	config  *config.Configuration
	service *service.Service
	router  *chi.Mux
}

func NewServer(config *config.Configuration, service *service.Service) *Server {
	return &Server{
		config:  config,
		service: service,
		router:  chi.NewRouter(),
	}
}

func (s *Server) Start(ctx context.Context) {
	s.routes()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.router,
	}

	slog.Warn("Server listing on port", "Port", s.config.Port)

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("server.Shutdown failed: ", "Error", err.Error())
		}
	})

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		slog.Error("http.ListenAndServe failed:", "Error", err.Error())
	}

	slog.Warn("Shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
