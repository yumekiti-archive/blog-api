package usecase

import (
	"github.com/yumekiti/blog-api/domain"
	"github.com/yumekiti/blog-api/domain/repository"
)

type CategoryUsecase interface {
	GetAll() ([]*domain.Category, error)
	Save(category *domain.Category) (*domain.Category, error)
	Remove(category *domain.Category) error
	Update(category *domain.Category) (*domain.Category, error)
}

type categoryUsecase struct {
	repository repository.CategoryRepository
}

func NewCategoryUsecase(repository repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		repository: repository,
	}
}

func (cu *categoryUsecase) GetAll() ([]*domain.Category, error) {
	foundCategorys, err := cu.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return foundCategorys, nil
}

func (cu *categoryUsecase) Save(category *domain.Category) (*domain.Category, error) {
	err := category.Set(category)
	if err != nil {
		return nil, err
	}

	savedCategory, err := cu.repository.Save(category)
	if err != nil {
		return nil, err
	}

	return savedCategory, nil
}

func (cu *categoryUsecase) Remove(category *domain.Category) error {
	err := cu.repository.Remove(category)
	if err != nil {
		return err
	}

	return nil
}

func (cu *categoryUsecase) Update(category *domain.Category) (*domain.Category, error) {
	err := category.Set(category)
	if err != nil {
		return nil, err
	}

	updatedCategory, err := cu.repository.Update(category)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}
