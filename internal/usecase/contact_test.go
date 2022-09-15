package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zorlovskiy/phonebook/internal/models"
	"github.com/zorlovskiy/phonebook/mock"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func contact(t *testing.T) *mock.MockContactRepository {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock.NewMockContactRepository(mockCtrl)

	//contact := usecase.NewContactUseCase(repo)

	return repo

}

func TestCreate(t *testing.T) {

	repo := contact(t)

	vas := &models.Contact{1, "vas", "kab", "5678"}

	repo.EXPECT().Create(vas).Return(nil).Times(1)

	fmt.Println("Test worked")

}

func TestGetByFName(t *testing.T) {

	repo := contact(t)

	var fname string = "vas"

	repo.EXPECT().GetByFName(fname).Return([]models.Contact{{1, "vas", "kab", "5678"}}, nil).Times(1)

	fmt.Println("Test worked")
}

func TestGetByNumber(t *testing.T) {

	repo := contact(t)

	var number string = "23523"

	repo.EXPECT().GetByNumber(number).Return([]models.Contact{{2, "nic", "vas", "23523"}}, nil).Times(1)

	fmt.Println("Test worked")
}

func TestUpdate(t *testing.T) {

	repo := contact(t)

	van := &models.Contact{4, "van", "kab", "2345"}

	repo.EXPECT().Update(van).Return(nil).Times(1)

	fmt.Println("Test worked")

}

func TestDelete(t *testing.T) {

	repo := contact(t)

	jhon := "4"

	repo.EXPECT().DeleteByID(jhon).Return(nil).Times(1)

	fmt.Println("Test worked")
}
