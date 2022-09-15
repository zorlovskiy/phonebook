package usecase

import "github.com/zorlovskiy/phonebook/internal/models"

type ContactUseCase struct {
	Repo ContactRepository
}

func NewContactUseCase(r ContactRepository) *ContactUseCase {
	return &ContactUseCase{
		Repo: r,
	}
}

func (uc ContactUseCase) Create(model *models.Contact) error {
	return uc.Repo.Create(model)
}

func (uc ContactUseCase) GetByNumber(number string) ([]models.Contact, error) {
	return uc.Repo.GetByNumber(number)
}

func (uc ContactUseCase) GetByFName(fname string) ([]models.Contact, error) {
	return uc.Repo.GetByFName(fname)
}

func (uc ContactUseCase) Update(model *models.Contact) error {
	return uc.Repo.Update(model)
}

func (uc ContactUseCase) Delete(ID string) error {
	return uc.Repo.DeleteByID(ID)
}
