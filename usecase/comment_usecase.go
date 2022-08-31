package usecase

import (
	"github.com/yumekiti/blog-api/domain"
	"github.com/yumekiti/blog-api/domain/repository"
)

type CommentUsecase interface {
	Save(comment *domain.Comment) (*domain.Comment, error)
}

type commentUsecase struct {
	repository repository.CommentRepository
}

func NewCommentUsecase(repository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{
		repository: repository,
	}
}

func (cu *commentUsecase) Save(comment *domain.Comment) (*domain.Comment, error) {
	err := comment.Set(comment)
	if err != nil {
		return nil, err
	}

	savedComment, err := cu.repository.Save(comment)
	if err != nil {
		return nil, err
	}

	return savedComment, nil
}
