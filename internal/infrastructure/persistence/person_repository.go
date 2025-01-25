package infrastructure

import (
	"fmt"
	"go-clean-architecture/internal/domain/entities"

	"gorm.io/gorm"
)

type PersonRepository interface {
	GetAll() ([]entities.Person, error)
	GetByID(id uint) (*entities.Person, error)
	Create(person entities.Person) error
	Update(person entities.Person) error
	Delete(id uint) error
}

type personRepositoryImpl struct {
	DB *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepositoryImpl {
	return &personRepositoryImpl{DB: db}
}

func (repo *personRepositoryImpl) GetAll() ([]entities.Person, error) {
	var persons []entities.Person
	err := repo.DB.Find(&persons).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(persons)
	return persons, nil
}

func (repo *personRepositoryImpl) GetByID(id uint) (*entities.Person, error) {
	var person entities.Person
	err := repo.DB.First(&person, id).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (repo *personRepositoryImpl) Create(person entities.Person) error {
	err := repo.DB.Create(&person).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *personRepositoryImpl) Update(person entities.Person) error {
	err := repo.DB.Save(&person).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *personRepositoryImpl) Delete(id uint) error {
	err := repo.DB.Delete(&entities.Person{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
