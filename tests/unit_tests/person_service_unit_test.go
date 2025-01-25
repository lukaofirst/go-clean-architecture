package unit_tests

import (
	"errors"
	"go-clean-architecture/internal/application/services"
	"go-clean-architecture/internal/domain/entities"
	"go-clean-architecture/tests/unit_tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonService_GetAll(t *testing.T) {
	mockRepo := new(mocks.MockPersonRepository)
	mockRepo.On("GetAll").Return([]entities.Person{{ID: 1, Name: "John"}}, nil)

	service := services.NewPersonService(mockRepo)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "John", result[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_GetByID(t *testing.T) {
	mockRepo := new(mocks.MockPersonRepository)
	mockRepo.On("GetByID", uint(1)).Return(&entities.Person{ID: 1, Name: "John"}, nil)

	service := services.NewPersonService(mockRepo)

	result, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "John", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_GetByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockPersonRepository)
	mockRepo.On("GetByID", uint(999)).Return((*entities.Person)(nil), errors.New("not found"))

	service := services.NewPersonService(mockRepo)

	result, err := service.GetByID(999)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestPersonService_Create(t *testing.T) {
	mockRepo := new(mocks.MockPersonRepository)
	newPerson := entities.Person{ID: 2, Name: "Jane"}
	mockRepo.On("Create", newPerson).Return(nil)

	service := services.NewPersonService(mockRepo)

	err := service.Create(newPerson)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_Update(t *testing.T) {
	mockRepo := new(mocks.MockPersonRepository)
	updatedPerson := entities.Person{ID: 1, Name: "John Updated"}
	mockRepo.On("Update", updatedPerson).Return(nil)

	service := services.NewPersonService(mockRepo)

	err := service.Update(updatedPerson)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_Delete(t *testing.T) {
	mockRepo := new(mocks.MockPersonRepository)
	mockRepo.On("Delete", uint(1)).Return(nil)

	service := services.NewPersonService(mockRepo)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
