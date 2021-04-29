package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ganganikalpana/covidLog/dto"
	"github.com/ganganikalpana/covidLog/service"
	"github.com/gorilla/mux"
)

type PersonHandler struct {
	service service.PersonService
}
type LogHandler struct {
	service service.LogService
}

func (p *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["person_id"]
	customer, err := p.service.GetPerson(id)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)

	}

}

func (p *PersonHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customer, err := p.service.GetAll()
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)

	}

}
func (p *PersonHandler) NewPerson(w http.ResponseWriter, r *http.Request) {

	var Request dto.NewPersonRequest
	err := json.NewDecoder(r.Body).Decode(&Request)
	if err != nil {
		writerResponse(w, http.StatusBadRequest, err.Error())
	} else {

		person, appError := p.service.NewPerson(Request)
		if appError != nil {
			writerResponse(w, appError.Code, appError.AsMessage())
		} else {
			writerResponse(w, http.StatusOK, person)

		}

	}
}
func (p *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["person_id"]
	var Request dto.NewPersonRequest
	err := json.NewDecoder(r.Body).Decode(&Request)
	if err != nil {
		writerResponse(w, http.StatusBadRequest, err.Error())
	} else {
		Request.PersonId = id
		person, appError := p.service.EditPerson(Request)
		if appError != nil {
			writerResponse(w, appError.Code, appError.AsMessage())
		} else {
			writerResponse(w, http.StatusOK, person)

		}

	}
}
func (p *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["person_id"]

	appError := p.service.DeletePerson(id)
	if appError != nil {
		writerResponse(w, appError.Code, appError.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, "deleted successfully")

	}

}
func (p *LogHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["company_id"]
	var Request dto.NewLogRequest
	err := json.NewDecoder(r.Body).Decode(&Request)
	if err != nil {
		writerResponse(w, http.StatusBadRequest, err.Error())
	} else {
		Request.CompanyId = id
		person, appError := p.service.CreateLog(Request)
		if appError != nil {
			writerResponse(w, appError.Code, appError.AsMessage())
		} else {
			writerResponse(w, http.StatusOK, person)

		}

	}
}
func (l *LogHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["person_id"]
	customer, err := l.service.FindById(id)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)

	}

}
func (l *LogHandler) FindLogByCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["company_id"]
	customer, err := l.service.FindLogByCompany(id)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)

	}

}
func (l *LogHandler) FindLogByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid := vars["company_id"]
	date:=vars["date"]
	customer, err := l.service.FindLogByDate(cid,date)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)

	}

}

func (p *PersonHandler) Person(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	vars := mux.Vars(r)
	id := vars["person_id"]
	customer, err := p.service.GetPerson(id)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)

	}

}

func writerResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
