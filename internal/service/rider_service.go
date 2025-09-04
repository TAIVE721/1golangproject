package service

import (
	"RiderApi/internal/domain"
	"RiderApi/internal/repository"
)

type RiderService interface {
	GetAllRiders() ([]domain.KamenRider, error)
}

type riderService struct {
	repo repository.RiderRepository
}

func NewRiderService(repo repository.RiderRepository) RiderService {
	return &riderService{
		repo: repo,
	}
}

func (s *riderService) GetAllRiders() ([]domain.KamenRider, error) {

	return s.repo.FindALL()
}
