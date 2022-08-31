package domain

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Tag) Set(tag *Tag) error {
	t.Body = tag.Body

	return nil
}
