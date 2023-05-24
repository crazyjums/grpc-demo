package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"server_streaming_pattern/common"
	pb "server_streaming_pattern/pb/genproto/product_info"
	"server_streaming_pattern/server/server"
)

const (
	PORT = ":50054"
)

func main() {
	lis, err := net.Listen("tcp", common.PortServerStreaming)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, server.NewServer())
	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
