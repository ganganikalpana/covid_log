package service

import (
	"fmt"

	"github.com/ganganikalpana/covidLog/domain"
	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
)

type CompanyService interface {
	GetAll() ([]domain.Company, *errs.AppError)
	GetCompany(id string) (*dto.CompanyResponse, *errs.AppError)
	NewCompany(req dto.NewCompanyRequest) (*dto.CompanyResponse, *errs.AppError)
}

type DefaultCompanyService struct {
	repo domain.CompanyRepository
}

func (d DefaultCompanyService) GetAll() ([]domain.Company, *errs.AppError) {
	fmt.Println("service")
	return d.repo.FindAll()
}

func (d DefaultCompanyService) GetCompany(id string) (*dto.CompanyResponse, *errs.AppError) {
	c, err := d.repo.FindById(id)
	if err != nil {
		return nil, err

	}
	response := c.ToDto()
	return &response, nil

}
func (d DefaultCompanyService) NewCompany(req dto.NewCompanyRequest) (*dto.CompanyResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Company{
		CompanyId:   req.CompanyId,
		CompanyName: req.CompanyName,
	}
	newCompany, err := d.repo.SaveCompany(a)
	if err != nil {
		return nil, err
	}
	response := newCompany.ToDto()
	return &response, nil

}

func NewCompanyService(repository domain.CompanyRepository) DefaultCompanyService {
	return DefaultCompanyService{repository}
}
