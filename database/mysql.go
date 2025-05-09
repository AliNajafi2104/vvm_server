package database

import (
	"github.com/AliNajafi2104/vvm_server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySqlDb() (Database[models.Product], error) {
	dsn := "root:Ckv64snt12345@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, db.Error
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	return &GormDatabase[models.Product]{DB: db}, nil
}
