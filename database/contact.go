package database

import (

	//"github.com/jinzhu/gorm"
	"github.com/zorlovskiy/phonebook/domain"
	"gorm.io/gorm"
)

type ContactStore struct {
	db *gorm.DB
}

func NewContactStore(db *gorm.DB) *ContactStore {
	return &ContactStore{
		db: db,
	}
}

//добавление контакта
func (c *ContactStore) Create(model *domain.Contact) error {
	return c.db.Create(model).Error

}

//поиск по имени
func (c *ContactStore) GetByFName(fName string) ([]domain.Contact, error) {
	var byName []domain.Contact
	err := c.db.Where("f_name LIKE ?", fName).Find(&byName).Error

	if err != nil {
		return nil, err
	}

	return byName, nil

}

//поиск по номеру
func (c *ContactStore) GetByNumber(number string) ([]domain.Contact, error) {
	var byNumber []domain.Contact
	err := c.db.Where("phone_number LIKE ?", number).Find(&byNumber).Error

	if err != nil {
		return nil, err
	}

	return byNumber, nil

}

//обновление контакта
func (c *ContactStore) Update(model *domain.Contact) error {
	return c.db.Save(model).Error
}

//удаление контакта
func (c *ContactStore) DeleteByID(ID string) error {
	var cntc []domain.Contact

	return c.db.Where("id = ?", ID).Delete(&cntc).Error

}
