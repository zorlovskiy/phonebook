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

func (c *ContactStore) Create(model *domain.Contact) error {
	return c.db.Create(model).Error

}

func (c *ContactStore) GetByFName(fName string) ([]domain.Contact, error) {
	var cntc []domain.Contact
	err := c.db.Where("f_name LIKE ?", fName).Find(&cntc).Error

	if err != nil {
		return nil, err
	}

	return cntc, nil

}

func (c *ContactStore) Update(model *domain.Contact) error {
	return c.db.Save(model).Error
}

func (c *ContactStore) DeleteByID(ID string) error {
	var cntc []domain.Contact

	return c.db.Where("id = ?", ID).Delete(&cntc).Error

}
