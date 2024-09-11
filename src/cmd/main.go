package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/danielfmpc/go_rpc_unary/src/pb/products"
	"github.com/danielfmpc/go_rpc_unary/src/repository"
	"google.golang.org/grpc"
)

type Server struct {
	products.ProductServiceServer
	productRepo *repository.ProductRepository
}

func (s *Server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.productRepo.Create(*product)
	if err != nil {
		return product, err
	}
	return &newProduct, nil
}

func (s *Server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	producList, err := s.productRepo.FindAll()
	if err != nil {
		return &producList, err
	}
	return &producList, nil
}

func main() {
	fmt.Println("Starting grpc server")

	server := Server{productRepo: &repository.ProductRepository{}}

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Erro on create listener: ", err)
	}

	sGrpc := grpc.NewServer()
	products.RegisterProductServiceServer(sGrpc, &server)

	if err := sGrpc.Serve(listener); err != nil {
		log.Fatalf("erro on server:", err)
	}

}
