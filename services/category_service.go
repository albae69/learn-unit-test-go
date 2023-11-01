package services

import (
	"errors"
	"learn-unit-test-go/entity"
	"learn-unit-test-go/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

// struct method for category repository service
func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}

}
