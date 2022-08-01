package domain

type Contact struct {
	Id          int    `json:"id"`
	FName       string `json:"fname"`
	Sname       string `json:"sname"`
	PhoneNumber int    `json:"number"`
}

/*
{   gorm:"primary_key"
	"fname": "Zaxar",
	"sname": "Orl",
	"number": "+79543678899"
}
*/
