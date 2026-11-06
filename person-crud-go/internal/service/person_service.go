package service

import (
	"github.com/example/person-crud-go/internal/entity"
	"github.com/example/person-crud-go/internal/repository"
)

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(repo *repository.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) FindAll() ([]entity.Person, error) {
	return s.repo.FindAll()
}

func (s *PersonService) FindByID(id uint) (*entity.Person, error) {
	return s.repo.FindByID(id)
}

func (s *PersonService) Save(person *entity.Person) error {
	return s.repo.Save(person)
}

func (s *PersonService) Delete(id uint) error {
	return s.repo.Delete(id)
}
