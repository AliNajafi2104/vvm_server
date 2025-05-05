package database

import (
	"log"

	"github.com/AliNajafi2104/vvm_server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqliteDb() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
