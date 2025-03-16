package config

import (
	"os"
	"strconv"
)

type Configuration struct {
	Port        uint
	DatabaseUrl string
}

func NewConfiguration() (*Configuration, error) {
	port, err := getPort()

	if err != nil {
		return nil, err
	}

	return &Configuration{
		Port:        port,
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}, nil
}

func getPort() (uint, error) {
	portStr := os.Getenv("PORT")
	portEnv, err := strconv.Atoi(portStr)

	if err != nil {
		return 0, err
	}

	return uint(portEnv), nil
}
