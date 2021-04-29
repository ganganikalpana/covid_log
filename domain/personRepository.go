package domain

import (
	"database/sql"
	"fmt"

	"github.com/ganganikalpana/covidLog/errs"
	"github.com/ganganikalpana/covidLog/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type PersonRepositoryDb struct {
	client *sqlx.DB
}

func (p PersonRepositoryDb) FindById(id string) (*Person, *errs.AppError) {
	fmt.Println("hello")

	byId := "select full_name,address,phone_number,person_id FROM person where person_id=?"
	var c Person
	err := p.client.Get(&c, byId, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("error while scanning customer table" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func (p PersonRepositoryDb) FindAll() ([]Person, *errs.AppError) {
	ALL := "SELECT * from person"
	c := make([]Person, 0)
	err := p.client.Select(&c, ALL)
	if err != nil {
		logger.Error("error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return c, nil
}

func (p PersonRepositoryDb) SavePerson(s Person) (*Person, *errs.AppError) {

	PersonIN, err := p.client.Exec("INSERT INTO person(full_name,person_id,address,phone_number) VALUES (?,?,?,?)", s.FullName, s.IdNumber, s.Address, s.PhoneNumber)
	if err != nil {
		logger.Error("Error while inserting person" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	id, err := PersonIN.LastInsertId()
	fmt.Println(id)
	if err != nil {
		logger.Error("Error while getting last insert id for new insertetd person" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	return &s, nil

}
func (p PersonRepositoryDb) UpdatePerson(s Person) (*Person, *errs.AppError) {

	sql := "UPDATE person SET full_name=?, address=?, phone_number=? WHERE person_id=?"
	result, err := p.client.Exec(sql, s.FullName, s.Address, s.PhoneNumber, s.IdNumber)
	if err != nil {
		logger.Error("error while updating person" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		logger.Error("Error while getting last updated id for updated person" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	return &s, nil
}
func (p PersonRepositoryDb) DeletePerson(id string) *errs.AppError {
	sql := "DELETE FROM person where person_id=?"
	result, err := p.client.Exec(sql, id)
	if err != nil {
		logger.Error("error while updating person" + err.Error())
		return errs.NewUnexpectedError("unexpected error from database")
	}
	i, err := result.LastInsertId()
	fmt.Println(i)
	if err != nil {
		logger.Error("Error while getting last updated id for updated person" + err.Error())
		return errs.NewUnexpectedError("unexpected error from database")
	}

	return nil
}

func NewPersonRepositoryDb(dbClient *sqlx.DB) PersonRepositoryDb {

	return PersonRepositoryDb{dbClient}

}
