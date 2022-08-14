package usecase

import "github.com/zorlovskiy/phonebook/internal/models"

type ContactUseCase struct {
	repo ContactRepository
}

func New(r ContactRepository) *ContactUseCase {
	return &ContactUseCase{
		repo: r,
	}
}

func (uc ContactUseCase) Create(model *models.Contact) error {
	return uc.repo.Create(model)

}

func (uc ContactUseCase) GetByFName(fName string) ([]models.Contact, error) {
	return uc.repo.GetByFName(fName)
}

func (uc ContactUseCase) GetByNumber(number string) ([]models.Contact, error) {
	return uc.repo.GetByNumber(number)
}

func (uc ContactUseCase) Update(model *models.Contact) error {
	return uc.repo.Update(model)
}

func (uc ContactUseCase) Delete(ID string) error {
	return uc.repo.Delete(ID)
}
