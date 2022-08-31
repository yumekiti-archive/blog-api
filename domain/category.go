package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Body string
}

func (t *Category) Set(category *Category) error {
	t.Body = category.Body

	return nil
}
