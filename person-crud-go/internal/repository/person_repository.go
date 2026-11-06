package repository

import (
	"github.com/example/person-crud-go/internal/entity"
	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) FindAll() ([]entity.Person, error) {
	var persons []entity.Person
	return persons, r.db.Find(&persons).Error
}

func (r *PersonRepository) FindByID(id uint) (*entity.Person, error) {
	var person entity.Person
	return &person, r.db.First(&person, id).Error
}

func (r *PersonRepository) Save(person *entity.Person) error {
	return r.db.Save(person).Error
}

func (r *PersonRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Person{}, id).Error
}
