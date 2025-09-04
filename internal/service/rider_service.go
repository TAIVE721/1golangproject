package service

import (
	"RiderApi/internal/domain"
	"RiderApi/internal/repository"
	"database/sql"
)

type RiderService interface {
	GetAllRiders() ([]domain.KamenRider, error)
}

type riderService struct {
	repo repository.RiderRepository
}

func NewRiderService(db *sql.DB) RiderService {
	return &riderService{
		repo: repository.NewRiderRepository(db),
	}
}

func (s *riderService) GetAllRiders() ([]domain.KamenRider, error) {

	return s.repo.FindALL()
}
