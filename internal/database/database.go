package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	db.AutoMigrate(&TestResult{})
	return db
}
