package database

import (
	"github.com/jswildcards/gotodo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gotodo.db"), &gorm.Config{})
	db.AutoMigrate(&models.Task{})
	return db, err
}
