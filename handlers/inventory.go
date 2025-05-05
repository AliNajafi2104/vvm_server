package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AliNajafi2104/vvm_server/models"
	"gorm.io/gorm"
)

type InventoryHandler struct {
	DB *gorm.DB
}

func (p *InventoryHandler) IncreaseProductCount(w http.ResponseWriter, req *http.Request) {

	count := req.URL.Query().Get("count")
	barcode := req.URL.Query().Get("barcode")
	var product models.Product

	p.DB.First(&product, barcode)

	num, err := strconv.Atoi(count)

	if err != nil {
		fmt.Println("Error:", err)
	}

	product.Count = num + product.Count

	p.DB.Save(&product)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (p *InventoryHandler) GetTotalInventoryValue(w http.ResponseWriter, req *http.Request) {

	var products []models.Product

	result := p.DB.Find(&products)
	if result.Error != nil {
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
