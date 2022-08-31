package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yumekiti/blog-api/domain"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Task{})

	return db
}
