package main

import (
	"fmt"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/database"
	"github.com/AliNajafi2104/vvm_server/handlers"
)

func main() {
	db, err := database.InitMySqlDb()
	if err != nil {
		fmt.Printf("error init db")
		return
	}

	productHandler := handlers.ProductHandler{DB: db}

	http.HandleFunc("/api/products", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			productHandler.GetProductByBarcode(w, req)

		case http.MethodPost:
			productHandler.CreateProduct(w, req)

		case http.MethodPut:
			productHandler.UpdateProduct(w, req)

		case http.MethodDelete:
			productHandler.DeleteProduct(w, req)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	})

	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)

}
