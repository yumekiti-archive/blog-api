package repository

import "github.com/yumekiti/blog-api/domain"

type CommentRepository interface {
	Save(comment *domain.Comment) (*domain.Comment, error)
}
