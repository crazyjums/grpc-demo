package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "unary_communication_pattern/order"
)

const (
	PORT = ":50052"
)

func main() {
	conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)
	productId, err := client.AddProduct(context.Background(), &pb.Product{
		Name:  "Product 1",
		Price: 100,
		Desc:  "Product 1 desc",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("ProductId : %v", productId)
	product, err := client.GetProduct(context.Background(), &pb.ProductId{Value: productId.Value})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("Product : %v", product)
}
