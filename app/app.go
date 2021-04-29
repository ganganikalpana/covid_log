package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ganganikalpana/covidLog/domain"
	"github.com/ganganikalpana/covidLog/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func SanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" {
		log.Fatal("SERVER_ADDRESS is  not defined")

	}
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("SERVER_PORT is not defined")
	}
	if os.Getenv("DB_USER") == "" {
		log.Fatal("DB_USER is not defined")
	}
	if os.Getenv("DB_PW") == "" {
		log.Fatal("DB_PW is not defined")
	}
	if os.Getenv("DB_ADDR") == "" {
		log.Fatal("DB_ADDR is not defined")
	}
	if os.Getenv("DB_NAME") == "" {
		log.Fatal("DB_NAME is not defined")
	}
	if os.Getenv("DB_PORT") == "" {
		log.Fatal("DB_PORT is not defined")
	}
}

func Start() {
	SanityCheck()
	router := mux.NewRouter()

	personRepositoryDb := domain.NewPersonRepositoryDb(getDbClient())
	companyRepositoryDb := domain.NewCompanyRepositoryDb(getDbClient())
	logRepositoryDb := domain.NewLogRepositoryDb(getDbClient())
	ph := PersonHandler{service.NewPersonService(personRepositoryDb)}
	lh := LogHandler{service.NewLogService(logRepositoryDb)}
	ch := CompanyHandler{service.NewCompanyService(companyRepositoryDb)}

	router.HandleFunc("/person/{person_id:[0-9]+}", ph.GetPerson).Methods(http.MethodGet)
	router.HandleFunc("/person", ph.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/person", ph.NewPerson).Methods(http.MethodPost)
	router.HandleFunc("/person/{person_id}", ph.UpdatePerson).Methods(http.MethodPatch)
	router.HandleFunc("/person/{person_id}", ph.DeletePerson).Methods(http.MethodDelete)

	router.HandleFunc("/company/{company_id}/log", lh.CreateLog).Methods(http.MethodPost)
	router.HandleFunc("/person/{person_id}/log", lh.FindById).Methods(http.MethodGet)
	router.HandleFunc("/company/{company_id}/log", lh.FindLogByCompany).Methods(http.MethodGet)
	router.HandleFunc("/company/{company_id}/{date}/log", lh.FindLogByDate).Methods(http.MethodGet)

	router.HandleFunc("/company/{company_id}", ch.GetCompany).Methods(http.MethodGet)
	router.HandleFunc("/company", ch.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/company", ch.NewCompany).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")

	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PW")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_Name")
	fmt.Println(dbName)

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPw, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(10)
	return client

}
