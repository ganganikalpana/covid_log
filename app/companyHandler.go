package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/service"
	"github.com/gorilla/mux"
)

type CompanyHandler struct {
	service service.CompanyService
}

func (c CompanyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("company")
	company, err := c.service.GetAll()
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, company)

	}

}

func (c CompanyHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["company_id"]
	company, err := c.service.GetCompany(id)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, company)

	}

}
func (c CompanyHandler) NewCompany(w http.ResponseWriter, r *http.Request) {
	var request dto.NewCompanyRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writerResponse(w, http.StatusBadRequest, err.Error())
	} else {
		company, appError := c.service.NewCompany(request)
		if appError != nil {
			writerResponse(w, appError.Code, appError.Message)
		} else {
			writerResponse(w, http.StatusCreated, company)
		}
	}
}
