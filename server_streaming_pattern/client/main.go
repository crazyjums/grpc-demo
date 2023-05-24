package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"server_streaming_pattern/common"
	pb "server_streaming_pattern/pb/genproto/product_info"
	"time"
)

func main() {
	conn, err := grpc.Dial(common.PortServerStreaming, grpc.WithInsecure())
	if err != nil {
		log.Fatalf(common.LogPrefixServerStreaming+"did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewProductInfoClient(conn)
	productId, err := client.AddProduct(context.Background(), &pb.Product{
		Name:  "apple11",
		Price: 100,
		Desc:  "this is a iphone11",
	})
	if err != nil {
		log.Fatalf(common.LogPrefixServerStreaming+"could not greet: %v", err)
		return
	}
	log.Printf(common.LogPrefixServerStreaming+"AddProduct one and returns ProductId : %v", productId)

	productId2, err := client.AddProduct(context.Background(), &pb.Product{
		Name:  "apple14",
		Price: 1000,
		Desc:  "this is a iphone14",
	})
	if err != nil {
		log.Fatalf(common.LogPrefixServerStreaming+"could not greet2: %v", err)
		return
	}
	log.Printf(common.LogPrefixServerStreaming+"AddProduct two and returns ProductId : %v", productId2)

	productId3, err := client.AddProduct(context.Background(), &pb.Product{
		Name:  "apple15",
		Price: 2000,
		Desc:  "this is a iphone15",
	})
	if err != nil {
		log.Fatalf(common.LogPrefixServerStreaming+"could not greet3: %v", err)
		return
	}
	log.Printf(common.LogPrefixServerStreaming+"AddProduct three and returns ProductId : %v", productId3)

	time.Sleep(time.Second)

	productStreams, err := client.SearchProducts(context.Background(), &pb.ProductName{Value: "apple"})
	if err != nil {
		log.Fatalf(common.LogPrefixServerStreaming+"could not greet: %v", err)
		return
	}
	i := 0
	for {
		product, err := productStreams.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(common.LogPrefixServerStreaming+"could not greet4: %v", err)
			return
		}
		i++
		log.Printf(common.LogPrefixServerStreaming+"server stream returns Product %d : %v", i, product)
	}
}
