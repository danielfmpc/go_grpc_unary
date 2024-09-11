package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danielfmpc/client/go_rpc_unary/src/pb/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Erro on get lcient", err)
	}

	defer conn.Close()

	client := products.NewProductServiceClient(conn)

	FindAll(client)
	Create(client)
	FindAll(client)
}

func Create(client products.ProductServiceClient) {
	newProduct := products.Product{
		Name:        "Monitor",
		Description: "Monitor LED Ultrawide 25'",
		Price:       1500,
		Quantity:    5,
	}
	producCreated, err := client.Create(context.Background(), &newProduct)

	if err != nil {
		log.Fatalf("Erro on create product ", err)
	}

	fmt.Printf("products: %+v\n", producCreated)
}

func FindAll(client products.ProductServiceClient) {
	productList, err := client.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatalf("Erro on list all product ", err)
	}

	fmt.Printf("products: %+v\n", productList)
}
