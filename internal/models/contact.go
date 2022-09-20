package models

type Contact struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	FName       string `json:"fname"`
	Sname       string `json:"sname"`
	PhoneNumber string `json:"number"`
}
