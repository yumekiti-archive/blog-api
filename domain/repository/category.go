package repository

import "github.com/yumekiti/blog-api/domain"

type CategoryRepository interface {
	GetAll() ([]*domain.Category, error)
	Save(category *domain.Category) (*domain.Category, error)
	Remove(category *domain.Category) error
	Update(category *domain.Category) (*domain.Category, error)
}
