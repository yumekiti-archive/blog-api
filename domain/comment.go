package domain

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Comment) Set(comment *Comment) error {
	t.Name = comment.Name
	t.Body = comment.Body

	return nil
}
