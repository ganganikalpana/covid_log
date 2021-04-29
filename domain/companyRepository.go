package domain

import (
	"database/sql"
	"fmt"

	"github.com/ganganikalpana/covidLog/errs"
	"github.com/ganganikalpana/covidLog/logger"
	"github.com/jmoiron/sqlx"
)

type CompanyRepositoryDb struct {
	client *sqlx.DB
}

func (c CompanyRepositoryDb) FindAll() ([]Company, *errs.AppError) {
	fmt.Println("query")
	ALL := "select * from company"
	p := make([]Company, 0)
	err := c.client.Select(&p, ALL)
	if err != nil {
		logger.Error("error while querying company table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return p, nil
}
func (c CompanyRepositoryDb) FindById(id string) (*Company, *errs.AppError) {

	byId := "select * FROM company where company_id=?"
	var p Company
	err := c.client.Get(&p, byId, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("company not found")
		} else {
			logger.Error("error while scanning company table" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &p, nil
}
func (c CompanyRepositoryDb) SaveCompany(a Company) (*Company, *errs.AppError) {
	insert := "INSERT INTO company(company_id,company_name) VALUES (?,?)"
	result, err := c.client.Exec(insert, a.CompanyId, a.CompanyName)
	if err != nil {
		logger.Error("error while saving a company" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	id, _ := result.LastInsertId()
	fmt.Println(id)
	return &a, nil

}
func NewCompanyRepositoryDb(dbClient *sqlx.DB) CompanyRepositoryDb {
	return CompanyRepositoryDb{dbClient}
}
