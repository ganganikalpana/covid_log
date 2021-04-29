package dto

import "github.com/ganganikalpana/covidLog/errs"

type NewPersonRequest struct {
	PersonName  string `json:"person_name"`
	Address     string `json:"address"`
	PersonId    string `json:"person_id"`
	PhoneNumber string `json:"phone_number"`
	CompanyId   string `json:"company_id"`
}

func (n NewPersonRequest) Validate() *errs.AppError {
	return nil

}
 