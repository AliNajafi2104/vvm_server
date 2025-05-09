// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/AliNajafi2104/vvm_server/database"
// 	"github.com/AliNajafi2104/vvm_server/handlers"
// 	"github.com/AliNajafi2104/vvm_server/middleware"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	db, err := database.InitMySqlDb()
// 	if err != nil {
// 		fmt.Printf("error init db")
// 		return
// 	}

// 	productHandler := handlers.ProductHandler{DB: db}
// 	inventoryHandler := handlers.InventoryHandler{DB: db}

// 	r := mux.NewRouter()

// 	r.HandleFunc("/login", handlers.Login).Methods("POST")

// 	protected := r.PathPrefix("/api").Subrouter()
// 	protected.Use(middleware.AuthMiddleware)

// 	protected.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
// 	protected.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
// 	protected.HandleFunc("/products/{barcode}", productHandler.GetProductByBarcode).Methods("GET")
// 	protected.HandleFunc("/products/{barcode}", productHandler.UpdateProduct).Methods("PATCH")
// 	protected.HandleFunc("/products/{barcode}", productHandler.DeleteProduct).Methods("DELETE")
// 	protected.HandleFunc("/inventory/{barcode}", inventoryHandler.IncreaseProductCount).Methods("POST")
// 	protected.HandleFunc("/inventory", inventoryHandler.GetTotalInventoryValue).Methods("GET")

// 	err = http.ListenAndServe(":8080", r)

// 	if err != nil {
// 		fmt.Printf("Error starting server: %v\n", err)
// 	} else {
// 		fmt.Println("Server is running!")
// 	}

// }
