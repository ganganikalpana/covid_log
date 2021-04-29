package service

import (
	"github.com/ganganikalpana/covidLog/domain"
	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
)

type PersonService interface {
	GetPerson(id string) (*dto.PersonResponse, *errs.AppError)
	GetAll() ([]domain.Person, *errs.AppError)
	NewPerson(req dto.NewPersonRequest) (*dto.PersonResponse, *errs.AppError)
	EditPerson(req dto.NewPersonRequest) (*dto.PersonResponse, *errs.AppError)
	DeletePerson(id string) *errs.AppError
}
type DefaultPersonService struct {
	repo domain.PersonRepository
}

func (s DefaultPersonService) GetPerson(id string) (*dto.PersonResponse, *errs.AppError) {

	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}
func (s DefaultPersonService) GetAll() ([]domain.Person, *errs.AppError) {
	return s.repo.FindAll()
}
func (s DefaultPersonService) NewPerson(c dto.NewPersonRequest) (*dto.PersonResponse, *errs.AppError) {
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Person{
		FullName:    c.PersonName,
		Address:     c.Address,
		IdNumber:    c.PersonId,
		PhoneNumber: c.PhoneNumber,
	}
	p, err := s.repo.SavePerson(a)
	if err != nil {
		return nil, err

	}
	response := p.ToDto()
	return &response, nil
}
func (s DefaultPersonService) EditPerson(c dto.NewPersonRequest) (*dto.PersonResponse, *errs.AppError) {
	a := domain.Person{
		FullName:    c.PersonName,
		Address:     c.Address,
		IdNumber:    c.PersonId,
		PhoneNumber: c.PhoneNumber,
	}
	p, err := s.repo.UpdatePerson(a)
	if err != nil {
		return nil, err

	}
	response := p.ToDto()
	return &response, nil

}
func (s DefaultPersonService) DeletePerson(id string) *errs.AppError {
	err := s.repo.DeletePerson(id)
	if err != nil {
		return err

	}
	return nil

}
func NewPersonService(repository domain.PersonRepository) DefaultPersonService {

	return DefaultPersonService{repository}
}
