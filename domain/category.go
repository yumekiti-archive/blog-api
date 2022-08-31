package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Body string
}
