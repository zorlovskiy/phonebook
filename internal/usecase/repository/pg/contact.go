package pg

import (
	"github.com/zorlovskiy/phonebook/internal/models"
	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		db: db,
	}
}

//добавление контакта
func (c *ContactRepository) Create(model *models.Contact) error {

	//id := model.Id

	//if id == "" {

	//	fmt.Println("Enter id")
	//	return nil

	//}
	return c.db.Create(model).Error

}

//поиск по имени
func (c *ContactRepository) GetByFName(fName string) ([]models.Contact, error) {
	var byName []models.Contact
	err := c.db.Where("f_name LIKE ?", fName).Find(&byName).Error

	if err != nil {
		return nil, err
	}

	return byName, nil

}

//поиск по номеру
func (c *ContactRepository) GetByNumber(number string) ([]models.Contact, error) {
	var byNumber []models.Contact
	err := c.db.Where("phone_number LIKE ?", number).Find(&byNumber).Error

	if err != nil {
		return nil, err
	}

	return byNumber, nil

}

//обновление контакта
func (c *ContactRepository) Update(model *models.Contact) error {
	return c.db.Save(model).Error
}

//удаление контакта
func (c *ContactRepository) DeleteByID(ID string) error {
	var cntc []models.Contact

	return c.db.Where("id = ?", ID).Delete(&cntc).Error

}
