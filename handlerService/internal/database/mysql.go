package database

import (
	"github.com/AliNajafi2104/vvm_server/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySqlDb() (*gorm.DB, error) {
	dsn := "root:Ckv64snt12345@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, db.Error
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Database interface {
	FindProductByBarcode(barcode string) (*models.Product, error)
	SaveProduct(product *models.Product) error
	CreateProduct(product *models.Product) error
	DeleteProduct(barcode string) error
}

type GormDatabase struct {
	DB *gorm.DB
}

func (g *GormDatabase) FindProductByBarcode(barcode string) (*models.Product, error) {
	var product models.Product
	result := g.DB.First(&product, "barcode = ?", barcode)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil

}
