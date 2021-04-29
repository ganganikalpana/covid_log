package dto

import "github.com/ganganikalpana/covidLog/errs"

type NewCompanyRequest struct {
	CompanyId   string `json:"company_id"`
	CompanyName string `json:"company_name"`
}

func (n NewCompanyRequest) Validate() *errs.AppError {
	return nil

}
