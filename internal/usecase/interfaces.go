package usecase

import "github.com/zorlovskiy/phonebook/internal/models"

type (
	Contact interface {
		Create(model *models.Contact) error
		GetByFName(fName string) ([]models.Contact, error)
		GetByNumber(number string) ([]models.Contact, error)
		Update(model *models.Contact) error
		Delete(ID string) error
	}

	ContactRepository interface {
		Create(model *models.Contact) error
		GetByFName(fName string) ([]models.Contact, error)
		GetByNumber(number string) ([]models.Contact, error)
		Update(model *models.Contact) error
		DeleteByID(ID string) error
	}
)
