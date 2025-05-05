package main

import (
	"fmt"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/database"
	"github.com/AliNajafi2104/vvm_server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitMySqlDb()
	if err != nil {
		fmt.Printf("error init db")
		return
	}

	productHandler := handlers.ProductHandler{DB: db}
	inventoryHandler := handlers.InventoryHandler{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{barcode}", productHandler.GetProductByBarcode).Methods("GET")
	r.HandleFunc("/api/products/{barcode}", productHandler.UpdateProduct).Methods("PATCH")
	r.HandleFunc("/api/products/{barcode}", productHandler.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/inventory/{barcode}", inventoryHandler.IncreaseProductCount).Methods("POST")

	fmt.Println("Server running...")
	err = http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
