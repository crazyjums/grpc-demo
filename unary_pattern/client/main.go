package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"unary_pattern/common"
	pb "unary_pattern/pb/genproto/product_info"
)

func main() {
	conn, err := grpc.Dial(common.PortUnary, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProductInfoClient(conn)
	productId, err := client.AddProduct(context.Background(), &pb.Product{
		Name:  "apple11",
		Price: 100,
		Desc:  "this is a iphone11",
	})
	if err != nil {
		log.Fatalf(common.LogPrefixUnary+"could not greet: %v", err)
		return
	}
	log.Printf(common.LogPrefixUnary+"ProductId : %v", productId)
	product, err := client.GetProduct(context.Background(), &pb.ProductId{Value: productId.Value})
	if err != nil {
		log.Fatalf(common.LogPrefixUnary+"could not greet: %v", err)
		return
	}
	log.Printf(common.LogPrefixUnary+"Product : %v", product)
}
