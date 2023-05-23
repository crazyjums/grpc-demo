package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"unary_communication_pattern/server/order"
	"unary_communication_pattern/server/server"
)

const (
	PORT = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	order.RegisterOrderManagementServer(s, &server.Server{})
	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
