package repository

import "learn-unit-test-go/entity"

// create an interface for contract
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
