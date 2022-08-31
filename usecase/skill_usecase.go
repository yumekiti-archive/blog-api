package usecase

import (
	"github.com/yumekiti/blog-api/domain"
	"github.com/yumekiti/blog-api/domain/repository"
)

type SkillUsecase interface {
	Get(id int) (*domain.Skill, error)
	GetAll() ([]*domain.Skill, error)
	Save(skill *domain.Skill) (*domain.Skill, error)
	Remove(skill *domain.Skill) error
	Update(skill *domain.Skill) (*domain.Skill, error)
}

type skillUsecase struct {
	skillRepo repository.SkillRepository
}

func NewSkillUsecase(skillRepo repository.SkillRepository) SkillUsecase {
	return &skillUsecase{skillRepo: skillRepo}
}

func (su *skillUsecase) Get(id int) (*domain.Skill, error) {
	foundSkill, err := su.skillRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return foundSkill, nil
}

func (su *skillUsecase) GetAll() ([]*domain.Skill, error) {
	foundSkill, err := su.skillRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return foundSkill, nil
}

func (su *skillUsecase) Save(skill *domain.Skill) (*domain.Skill, error) {
	err := skill.Set(skill)
	if err != nil {
		return nil, err
	}

	savedSkill, err := su.skillRepo.Save(skill)
	if err != nil {
		return nil, err
	}

	return savedSkill, nil
}

func (su *skillUsecase) Remove(skill *domain.Skill) error {
	return su.skillRepo.Remove(skill)
}

func (su *skillUsecase) Update(skill *domain.Skill) (*domain.Skill, error) {
	err := skill.Set(skill)
	if err != nil {
		return nil, err
	}

	updatedSkill, err := su.skillRepo.Update(skill)
	if err != nil {
		return nil, err
	}

	return updatedSkill, nil
}
