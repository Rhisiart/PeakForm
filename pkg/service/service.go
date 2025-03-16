package service

import "github.com/Rhisiart/PeakForm/pkg/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
