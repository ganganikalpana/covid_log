package service

import (
	"time"

	"github.com/ganganikalpana/covidLog/domain"
	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/errs"
)

type LogService interface {
	CreateLog(s dto.NewLogRequest) (*dto.NewLogResponse, *errs.AppError)
	FindById(id string) ([]dto.NewLogResponses, *errs.AppError)
	FindLogByCompany(id string) ([]dto.NewLogResponses, *errs.AppError)
	FindLogByDate(id string, date string) ([]dto.NewLogResponses, *errs.AppError)
}
type DefaultLogService struct {
	repo domain.LogRepository
}

func (d DefaultLogService) CreateLog(s dto.NewLogRequest) (*dto.NewLogResponse, *errs.AppError) {

	a := domain.Log{
		LogId:       "",
		CompanyId:   s.CompanyId,
		PersonId:    s.PersonId,
		DateAndTime: time.Now().Format("2006-01-02 15:04:05"),
		LogDate:     time.Now().Format("2006-01-02"),
	}

	p, err := d.repo.CreateLog(a)
	if err != nil {
		return nil, err

	}
	response := p.ToDto()
	return &response, nil
}

func (d DefaultLogService) FindById(id string) ([]dto.NewLogResponses, *errs.AppError) {
	c, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return c, nil

}
func (d DefaultLogService) FindLogByCompany(id string) ([]dto.NewLogResponses, *errs.AppError) {
	c, err := d.repo.FindLogByCompany(id)
	if err != nil {
		return nil, err
	}
	return c, nil

}
func (d DefaultLogService) FindLogByDate(id, date string) ([]dto.NewLogResponses, *errs.AppError) {
	c, err := d.repo.FindLogByDate(id, date)
	if err != nil {
		return nil, err
	}
	return c, nil

}
func NewLogService(repository domain.LogRepository) DefaultLogService {

	return DefaultLogService{repository}
}
