package main

import (
	"context"
	"log"
	"net"

	pb "github.com/AliNajafi2104/vvm_server/common/proto"
	"github.com/AliNajafi2104/vvm_server/handlerService/internal/database"
	"github.com/AliNajafi2104/vvm_server/handlerService/internal/handlers"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHandlersServiceServer
	ProductHandler *handlers.ProductHandler
}

func (s *server) GetProductByBarcode(ctx context.Context, req *pb.Barcode) (*pb.Product, error) {
	// // Implement your logic here
	product, err := s.ProductHandler.FetchProductByBarcode(req.Barcode)
	if err != nil {
		return nil, err
	}
	return &pb.Product{
		Name:    product.Name,
		Barcode: product.Barcode,
		Price:   product.Price,
		Count:   int32(product.Count),
	}, nil

}

func main() {

	db, err := database.InitMySqlDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	productHandler := &handlers.ProductHandler{
		DB: db,
	}

	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)

	}

	grpcServer := grpc.NewServer()

	pb.RegisterHandlersServiceServer(grpcServer, &server{ProductHandler: productHandler})

	log.Printf("Server is listening on port %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
