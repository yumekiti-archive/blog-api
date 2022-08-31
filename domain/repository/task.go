package repository

import (
	"github.com/yumekiti/blog-api/domain"
)

type TaskRepository interface {
	Get(id int) (*domain.Task, error)
	GetAll() ([]*domain.Task, error)
	Save(task *domain.Task) (*domain.Task, error)
	Remove(task *domain.Task) error
	Update(task *domain.Task) (*domain.Task, error)
}
