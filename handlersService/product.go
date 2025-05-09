package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/models"
	"github.com/gorilla/mux"
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

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {

	var product models.Product

	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result := p.DB.Create(&product)
	if result.Error != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, req *http.Request) {

	var updatedProduct models.Product

	if err := json.NewDecoder(req.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	barcode := vars["barcode"]
	var product models.Product

	p.DB.First(&product, barcode)

	product.Count = updatedProduct.Count
	product.Price = updatedProduct.Price
	product.Name = updatedProduct.Name
	p.DB.Save(&product)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	barcode := vars["barcode"]
	p.DB.Delete(&models.Product{}, barcode)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(barcode)

}

func (p *ProductHandler) GetAllProducts(w http.ResponseWriter, req *http.Request) {

	var products []models.Product

	result := p.DB.Find(&products)
	if result.Error != nil {
		http.Error(w, "error getting products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
