package dto

type NewLogRequest struct {
	PersonId  string `json:"person_id"`
	CompanyId string
}

type NewLogResponse struct {
	LogId       string 
	DateAndTime string
	LogDate     string
}
type NewLogResponses struct{
	LogId       string `db:"log_id"`
	DateAndTime string `db:"date_and_time"`
	LogDate     string `db:"logDate"`
	PersonId    string `db:"person_id"`
	CompanyId   string `db:"company_id"`
	PersonName string `db:"full_name"`
	Address string `db:"address"`
	PhoneNumber string `db:"phone_number"`

}
