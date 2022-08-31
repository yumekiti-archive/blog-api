package repository

import "github.com/yumekiti/blog-api/domain"

type SkillRepository interface {
	Get(id int) (*domain.Skill, error)
	GetAll() ([]*domain.Skill, error)
	Save(skill *domain.Skill) (*domain.Skill, error)
	Remove(skill *domain.Skill) error
	Update(skill *domain.Skill) (*domain.Skill, error)
}
