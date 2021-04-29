package domain

import (
	"fmt"
	"strconv"

	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
	"github.com/ganganikalpana/covidLog/logger"
	"github.com/jmoiron/sqlx"
)

type LogRepositoryDb struct {
	client *sqlx.DB
}

func (p LogRepositoryDb) CreateLog(s Log) (*Log, *errs.AppError) {
	log, err := p.client.Exec("INSERT INTO log(company_id,person_id,date_and_time,logDate) VALUES (?,?,?,?)", s.CompanyId, s.PersonId, s.DateAndTime, s.LogDate)
	if err != nil {
		logger.Error("Error while recording a log" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	id, err := log.LastInsertId()
	fmt.Println(id)
	if err != nil {
		logger.Error("Error while getting last insert id for new insertetd person" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	s.LogId = strconv.FormatInt(id, 10)
	return &s, nil

}
func (l LogRepositoryDb) FindById(id string) ([]dto.NewLogResponses, *errs.AppError) {
	sql := "select p.person_id,p.full_name,l.company_id,p.address,p.phone_number,date_and_time,logDate,log_id FROM log l inner join person p on l.person_id=p.person_id where l.person_id=?"
	c := make([]dto.NewLogResponses, 0)
	err := l.client.Select(&c, sql, id)

	if err != nil {
		logger.Error("error while retrieving logs" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	return c, nil

}
func (l LogRepositoryDb) FindLogByCompany(id string) ([]dto.NewLogResponses, *errs.AppError) {
	sql := "select p.person_id,p.full_name,l.company_id,p.address,p.phone_number,date_and_time,logDate,log_id FROM log l inner join person p on l.person_id=p.person_id where l.company_id=?"
	c := make([]dto.NewLogResponses, 0)
	err := l.client.Select(&c, sql, id)

	if err != nil {
		logger.Error("error while retrieving logs" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	return c, nil

}
func (l LogRepositoryDb) FindLogByDate(id string, day string) ([]dto.NewLogResponses, *errs.AppError) {
	sql := "select p.person_id,p.full_name,l.company_id,p.address,p.phone_number,date_and_time,logDate,log_id FROM log l inner join person p on l.person_id=p.person_id where l.logDate=? and l.company_id=?"
	c := make([]dto.NewLogResponses, 0)
	err := l.client.Select(&c, sql, day, id)

	if err != nil {
		logger.Error("error while retrieving logs" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	return c, nil

}
func NewLogRepositoryDb(dbClient *sqlx.DB) LogRepositoryDb {

	return LogRepositoryDb{dbClient}

}
