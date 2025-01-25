package mocks

import (
	"go-clean-architecture/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockPersonRepository struct {
	mock.Mock
}

func (m *MockPersonRepository) GetAll() ([]entities.Person, error) {
	args := m.Called()
	return args.Get(0).([]entities.Person), args.Error(1)
}

func (m *MockPersonRepository) GetByID(id uint) (*entities.Person, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Person), args.Error(1)
}

func (m *MockPersonRepository) Create(person entities.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockPersonRepository) Update(person entities.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockPersonRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
