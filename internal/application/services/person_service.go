package services

import (
	"go-clean-architecture/internal/domain/entities"
	infrastructure "go-clean-architecture/internal/infrastructure/persistence"
)

type PersonService interface {
	GetAll() ([]entities.Person, error)
	GetByID(id uint) (*entities.Person, error)
	Create(person entities.Person) error
	Update(person entities.Person) error
	Delete(id uint) error
}

type personService struct {
	Repo infrastructure.PersonRepository
}

func NewPersonService(repo infrastructure.PersonRepository) PersonService {
	return &personService{Repo: repo}
}

func (service *personService) GetAll() ([]entities.Person, error) {
	return service.Repo.GetAll()
}

func (service *personService) GetByID(id uint) (*entities.Person, error) {
	return service.Repo.GetByID(id)
}

func (service *personService) Create(person entities.Person) error {
	return service.Repo.Create(person)
}

func (service *personService) Update(person entities.Person) error {
	return service.Repo.Update(person)
}

func (service *personService) Delete(id uint) error {
	return service.Repo.Delete(id)
}
