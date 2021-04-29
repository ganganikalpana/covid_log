package domain

import (
	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
)

type Log struct {
	CompanyId   string `db:"company_id"`
	LogId       string `db:"log_id"`
	DateAndTime string `db:"date_and_time"`
	LogDate     string `db:"logDate"`
	PersonId    string `db:"person_id"`
}

type LogRepository interface {
	CreateLog(L Log) (*Log, *errs.AppError)
	FindById(id string) ([]dto.NewLogResponses, *errs.AppError)
	FindLogByCompany(id string) ([]dto.NewLogResponses, *errs.AppError)
	FindLogByDate(id string, day string) ([]dto.NewLogResponses, *errs.AppError)
}

func (l Log) ToDto() dto.NewLogResponse {
	return dto.NewLogResponse{
		LogId:       l.LogId,
		DateAndTime: l.DateAndTime,
	}

}
