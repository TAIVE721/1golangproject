package services

import (
	"RiderApi/internal/domain"
	"RiderApi/internal/repositories"
)

type RiderService interface {
	GetAll() ([]domain.KamenRider, error)
	GetById(id int) (domain.KamenRider, error)
	Post(rider domain.KamenRider) (domain.KamenRider, error)
	Patch(rider domain.KamenRider, id int) (domain.KamenRider, error)
	Delete(id int) (int, error)
}

type riderService struct {
	repository repositories.RiderRepository
}

func NewRiderService(r repositories.RiderRepository) RiderService {
	return &riderService{
		repository: r,
	}
}

func (s *riderService) GetAll() ([]domain.KamenRider, error) {

	return s.repository.GetAll()

}

func (s *riderService) GetById(id int) (domain.KamenRider, error) {

	return s.repository.GetById(id)
}

func (s *riderService) Post(rider domain.KamenRider) (domain.KamenRider, error) {

	return s.repository.Post(rider)

}
func (s *riderService) Patch(rider domain.KamenRider, id int) (domain.KamenRider, error) {

	return s.repository.Patch(rider, id)

}
func (s *riderService) Delete(id int) (int, error) {

	return s.repository.Delete(id)

}
