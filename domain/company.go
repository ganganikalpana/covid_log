package domain

import (
	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
)

type Company struct {
	CompanyId   string `db:"company_id"`
	CompanyName string `db:"company_name"`
}
type CompanyRepository interface {
	FindAll() ([]Company, *errs.AppError)
	FindById(id string) (*Company, *errs.AppError)
	SaveCompany(c Company) (*Company, *errs.AppError)
}

func (c Company) ToDto() dto.CompanyResponse {
	return dto.CompanyResponse{
		CompanyId:   c.CompanyId,
		CompanyName: c.CompanyName,
	}
}
