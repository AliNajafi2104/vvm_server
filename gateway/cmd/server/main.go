package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/AliNajafi2104/vvm_server/common/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to the grpc server %v", err)
	}

	defer conn.Close()

	r := mux.NewRouter()
	client := pb.NewHandlersServiceClient(conn)

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

		defer cancel()

		req := &pb.Barcode{Barcode: "123"}

		resp, err := client.GetProductByBarcode(ctx, req)
		if err != nil {
			http.Error(w, "error calling grpc server", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	}).Methods("GET")

	fmt.Println("HTTP server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
