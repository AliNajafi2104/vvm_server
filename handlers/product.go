package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/database"
	"github.com/AliNajafi2104/vvm_server/models"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	DB database.Database[models.Product]
}

type ProductHTTPHandler interface {
	GetProductByBarcode(w http.ResponseWriter, req *http.Request)
	CreateProduct(w http.ResponseWriter, req *http.Request)
	UpdateProduct(w http.ResponseWriter, req *http.Request)
	DeleteProduct(w http.ResponseWriter, req *http.Request)
	GetAllProducts(w http.ResponseWriter, req *http.Request)
}

func (p *ProductHandler) GetProductByBarcode(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	barcode := vars["barcode"]

	if barcode == "" {
		http.Error(w, "missing barcode param", http.StatusBadRequest)
		return
	}

	product, err := p.DB.FindByID(barcode)

	if err != nil {
		http.Error(w, "could not find product", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)

}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {

	var product models.Product

	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := p.DB.CreateEntity(&product)
	if err != nil {
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

	err := p.DB.UpdateEntity(&updatedProduct)
	if err != nil {
		http.Error(w, "error updating product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedProduct)
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	barcode := vars["barcode"]

	err := p.DB.DeleteByID(barcode)
	if err != nil {
		http.Error(w, "error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(barcode)

}

func (p *ProductHandler) GetAllProducts(w http.ResponseWriter, req *http.Request) {

	products, err := p.DB.FindAll()
	if err != nil {
		http.Error(w, "error getting products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}
