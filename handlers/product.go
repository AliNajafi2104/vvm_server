package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (p *ProductHandler) GetProductByBarcode(w http.ResponseWriter, req *http.Request) {

	barcode := req.URL.Query().Get("barcode")

	if barcode == "" {
		http.Error(w, "missing barcode param", http.StatusBadRequest)
		return
	}

	var product models.Product

	result := p.DB.Where("barcode = ?", barcode).First((&product))

	if result.Error != nil {
		http.Error(w, "error getting product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)

}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {

}
