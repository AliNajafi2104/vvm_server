package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliNajafi2104/vvm_server/database"
	"github.com/AliNajafi2104/vvm_server/models"
)

type InventoryHandler struct {
	DB database.Database[models.Product]
}

type InventoryHTTPHandler interface {
	IncreaseProductCount(w http.ResponseWriter, req *http.Request)
	GetTotalInventoryValue(w http.ResponseWriter, req *http.Request)
}

func (p *InventoryHandler) IncreaseProductCount(w http.ResponseWriter, req *http.Request) {

	count := req.URL.Query().Get("count")
	barcode := req.URL.Query().Get("barcode")

	product, err := p.DB.FindByID(barcode)

	num, err := strconv.Atoi(count)

	if err != nil {
		fmt.Println("Error:", err)
	}
	product.Count = num + product.Count

	p.DB.UpdateEntity(product)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (p *InventoryHandler) GetTotalInventoryValue(w http.ResponseWriter, req *http.Request) {

	products, err := p.DB.FindAll()

	if err != nil {
		http.Error(w, "error getting products", http.StatusInternalServerError)
		return
	}

	var sum float64

	for _, product := range products {
		sum += product.Price * float64(product.Count)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sum)

}
