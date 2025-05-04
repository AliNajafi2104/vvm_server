package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/models"
)

type ProductHandler struct {
	DB *sql.DB
}

func (p *ProductHandler) GetProductByBarcode(w http.ResponseWriter, req *http.Request) {

	product := models.Product{
		ID:      "1233434",
		Price:   20,
		Name:    "testproduct",
		Barcode: "123",
		Count:   1,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)

}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {

}
