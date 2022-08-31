package repository

import "github.com/yumekiti/blog-api/domain"

type TagRepository interface {
	GetAll() ([]*domain.Tag, error)
	Save(tag *domain.Tag) (*domain.Tag, error)
	Remove(tag *domain.Tag) error
	Update(tag *domain.Tag) (*domain.Tag, error)
}
