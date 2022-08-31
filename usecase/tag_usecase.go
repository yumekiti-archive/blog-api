package usecase

import (
	"github.com/yumekiti/blog-api/domain"
	"github.com/yumekiti/blog-api/domain/repository"
)

type TagUsecase interface {
	GetAll() ([]*domain.Tag, error)
	Save(tag *domain.Tag) (*domain.Tag, error)
	Remove(tag *domain.Tag) error
	Update(tag *domain.Tag) (*domain.Tag, error)
}

type tagUsecase struct {
	tagRepo repository.TagRepository
}

func NewTagUsecase(tagRepo repository.TagRepository) TagUsecase {
	return &tagUsecase{tagRepo: tagRepo}
}

func (tu *tagUsecase) GetAll() ([]*domain.Tag, error) {
	foundTag, err := tu.tagRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return foundTag, nil
}

func (tu *tagUsecase) Save(tag *domain.Tag) (*domain.Tag, error) {
	err := tag.Set(tag)
	if err != nil {
		return nil, err
	}

	savedTag, err := tu.tagRepo.Save(tag)
	if err != nil {
		return nil, err
	}

	return savedTag, nil
}

func (tu *tagUsecase) Remove(tag *domain.Tag) error {
	err := tu.tagRepo.Remove(tag)
	if err != nil {
		return err
	}

	return nil
}

func (tu *tagUsecase) Update(tag *domain.Tag) (*domain.Tag, error) {
	err := tag.Set(tag)
	if err != nil {
		return nil, err
	}

	updatedTag, err := tu.tagRepo.Update(tag)
	if err != nil {
		return nil, err
	}

	return updatedTag, nil
}
