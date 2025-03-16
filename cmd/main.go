package main

import (
	"context"
	"log/slog"

	"github.com/Rhisiart/PeakForm/pkg/api"
	"github.com/Rhisiart/PeakForm/pkg/config"
	"github.com/Rhisiart/PeakForm/pkg/repository"
	"github.com/Rhisiart/PeakForm/pkg/service"
)

func main() {
	config, err := config.NewConfiguration()

	if err != nil {
		slog.Error("Configuration missing", "Error", err.Error())
		return
	}

	db := repository.NewDatabase(config.DatabaseUrl)
	err = db.Connect()
	defer db.Close()

	if err != nil {
		slog.Error("Unable to connect to the database", "Error", err.Error())
		return
	}

	repo := repository.NewRepository(db.Db)
	service := service.NewService(repo)

	server := api.NewServer(config, service)
	ctx := context.Background()

	server.Start(ctx)
}
