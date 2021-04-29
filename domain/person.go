package domain

import (
	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
)

type Person struct {
	FullName    string `db:"full_name"`
	IdNumber    string `db:"person_id"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
}

type PersonRepository interface {
	FindById(id string) (*Person, *errs.AppError)
	FindAll() ([]Person, *errs.AppError)
	SavePerson(c Person) (*Person, *errs.AppError)
	UpdatePerson(c Person) (*Person, *errs.AppError)
	DeletePerson(id string) *errs.AppError
}

func (p Person) ToDto() dto.PersonResponse {
	return dto.PersonResponse{
		FullName:    p.FullName,
		PersonId:    p.IdNumber,
		Address:     p.Address,
		PhoneNumber: p.PhoneNumber,
	}

}
