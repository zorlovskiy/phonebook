package models

type Contact struct {
	Id          string `json:"id, primary_key"`
	FName       string `json:"fname"`
	Sname       string `json:"sname"`
	PhoneNumber string `json:"number"`
}
