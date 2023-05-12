package repository

import "github.com/dptsi/backend-yogameleniawan/03-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
