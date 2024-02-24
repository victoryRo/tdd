package handler

import (
	"errors"
	"github.com/victoryRo/tdd/class5/api/model"
)

type StorageMockError struct{}

func (sme *StorageMockError) Create(person *model.Person) error {
	return errors.New("error mock")
}

func (sme *StorageMockError) Update(ID int, person *model.Person) error {
	return errors.New("error mock")
}

func (sme *StorageMockError) Delete(ID int) error {
	return errors.New("error mock")
}

func (sme *StorageMockError) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("error mock")
}

func (sme *StorageMockError) GetAll() (model.Persons, error) {
	return model.Persons{}, errors.New("error mock")
}

type StorageMockOK struct{}

func (smo *StorageMockOK) Create(person *model.Person) error {
	return nil
}

func (smo *StorageMockOK) Update(ID int, person *model.Person) error {
	return nil
}

func (smo *StorageMockOK) Delete(ID int) error {
	return nil
}

func (smo *StorageMockOK) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}

func (smo *StorageMockOK) GetAll() (model.Persons, error) {
	return model.Persons{}, nil
}
