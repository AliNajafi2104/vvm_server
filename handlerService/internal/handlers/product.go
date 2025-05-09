package handlers

import (
	"errors"

	"github.com/AliNajafi2104/vvm_server/common/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (p *ProductHandler) FetchProductByBarcode(barcode string) (*models.Product, error) {

	if barcode == "" {
		return nil, errors.New("missing barcode param")

	}

	var product models.Product
	result := p.DB.First(&product, barcode)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}

	return &product, nil

}

func (p *ProductHandler) CreateProduct(product *models.Product) error {

	result := p.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductHandler) UpdateProduct(product models.Product) error {

	var origProduct models.Product
	p.DB.First(&origProduct, product.Barcode)

	origProduct.Count = product.Count
	origProduct.Name = product.Name
	origProduct.Price = product.Price

	p.DB.Save(&product)

	return nil
}

func (p *ProductHandler) DeleteProduct(barcode string) error {

	p.DB.Delete(&models.Product{}, barcode)

	return nil
}

func (p *ProductHandler) GetAllProducts() ([]models.Product, error) {

	var products []models.Product

	result := p.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil

}
