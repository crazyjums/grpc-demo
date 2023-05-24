package main

import (
	"client_streaming_pattern/common"
	pb "client_streaming_pattern/pb/genproto/product_info"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(common.PortClientStreaming, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(common.LogPrefixClientStreaming+"did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProductInfoClient(conn)
	productId, err := client.AddProduct(context.Background(), &pb.Product{
		Name:  "apple11",
		Price: 100,
		Desc:  "this is a iphone11",
	})
	if err != nil {
		log.Fatalf(common.LogPrefixClientStreaming+"could not greet: %v", err)
		return
	}
	log.Printf(common.LogPrefixClientStreaming+"AddProduct one and returns ProductId : %v", productId)

	product2 := &pb.Product{
		Name:  "apple14",
		Price: 1000,
		Desc:  "this is a iphone14",
	}
	productId2, err := client.AddProduct(context.Background(), product2)
	if err != nil {
		log.Fatalf(common.LogPrefixClientStreaming+"could not greet2: %v", err)
		return
	}
	log.Printf(common.LogPrefixClientStreaming+"AddProduct two and returns ProductId : %v", productId2)

	product3 := &pb.Product{
		Name:  "apple15",
		Price: 2000,
		Desc:  "this is a iphone15",
	}
	productId3, err := client.AddProduct(context.Background(), product3)
	if err != nil {
		log.Fatalf(common.LogPrefixClientStreaming+"could not greet3: %v", err)
		return
	}
	log.Printf(common.LogPrefixClientStreaming+"AddProduct three and returns ProductId : %v", productId3)

	time.Sleep(time.Second)

	stream, err := client.UpdateProducts(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}

	time.Sleep(time.Second)
	if err := stream.Send(product2); err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	time.Sleep(time.Second)
	if err := stream.Send(product3); err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
