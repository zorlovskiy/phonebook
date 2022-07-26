// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/zorlovskiy/phonebook/internal/models"
)

// MockContact is a mock of Contact interface.
type MockContact struct {
	ctrl     *gomock.Controller
	recorder *MockContactMockRecorder
}

// MockContactMockRecorder is the mock recorder for MockContact.
type MockContactMockRecorder struct {
	mock *MockContact
}

// NewMockContact creates a new mock instance.
func NewMockContact(ctrl *gomock.Controller) *MockContact {
	mock := &MockContact{ctrl: ctrl}
	mock.recorder = &MockContactMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContact) EXPECT() *MockContactMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockContact) Create(model *models.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockContactMockRecorder) Create(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContact)(nil).Create), model)
}

// Delete mocks base method.
func (m *MockContact) Delete(ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockContactMockRecorder) Delete(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContact)(nil).Delete), ID)
}

// GetByFName mocks base method.
func (m *MockContact) GetByFName(fName string) ([]models.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByFName", fName)
	ret0, _ := ret[0].([]models.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFName indicates an expected call of GetByFName.
func (mr *MockContactMockRecorder) GetByFName(fName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFName", reflect.TypeOf((*MockContact)(nil).GetByFName), fName)
}

// GetByNumber mocks base method.
func (m *MockContact) GetByNumber(number string) ([]models.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNumber", number)
	ret0, _ := ret[0].([]models.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNumber indicates an expected call of GetByNumber.
func (mr *MockContactMockRecorder) GetByNumber(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNumber", reflect.TypeOf((*MockContact)(nil).GetByNumber), number)
}

// Update mocks base method.
func (m *MockContact) Update(model *models.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockContactMockRecorder) Update(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockContact)(nil).Update), model)
}

// MockContactRepository is a mock of ContactRepository interface.
type MockContactRepository struct {
	ctrl     *gomock.Controller
	recorder *MockContactRepositoryMockRecorder
}

// MockContactRepositoryMockRecorder is the mock recorder for MockContactRepository.
type MockContactRepositoryMockRecorder struct {
	mock *MockContactRepository
}

// NewMockContactRepository creates a new mock instance.
func NewMockContactRepository(ctrl *gomock.Controller) *MockContactRepository {
	mock := &MockContactRepository{ctrl: ctrl}
	mock.recorder = &MockContactRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContactRepository) EXPECT() *MockContactRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockContactRepository) Create(model *models.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockContactRepositoryMockRecorder) Create(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContactRepository)(nil).Create), model)
}

// DeleteByID mocks base method.
func (m *MockContactRepository) DeleteByID(ID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockContactRepositoryMockRecorder) DeleteByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockContactRepository)(nil).DeleteByID), ID)
}

// GetByFName mocks base method.
func (m *MockContactRepository) GetByFName(fName string) ([]models.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByFName", fName)
	ret0, _ := ret[0].([]models.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFName indicates an expected call of GetByFName.
func (mr *MockContactRepositoryMockRecorder) GetByFName(fName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFName", reflect.TypeOf((*MockContactRepository)(nil).GetByFName), fName)
}

// GetByNumber mocks base method.
func (m *MockContactRepository) GetByNumber(number string) ([]models.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNumber", number)
	ret0, _ := ret[0].([]models.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNumber indicates an expected call of GetByNumber.
func (mr *MockContactRepositoryMockRecorder) GetByNumber(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNumber", reflect.TypeOf((*MockContactRepository)(nil).GetByNumber), number)
}

// Update mocks base method.
func (m *MockContactRepository) Update(model *models.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockContactRepositoryMockRecorder) Update(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockContactRepository)(nil).Update), model)
}
