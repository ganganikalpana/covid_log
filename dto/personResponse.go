package dto

type PersonResponse struct{
	PersonId string `json:"nic_number"`
	FullName string `json:"Name"`
	Address string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}