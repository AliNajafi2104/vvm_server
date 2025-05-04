package main

import (
	"fmt"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/database"
	"github.com/AliNajafi2104/vvm_server/handlers"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		return
	}

	productHandler := handlers.ProductHandler{DB: db}

	http.HandleFunc("/api/products", productHandler.GetProductByBarcode)

	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)

	defer db.Close()
}
