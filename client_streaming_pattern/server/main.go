package main

import (
	"client_streaming_pattern/common"
	pb "client_streaming_pattern/pb/genproto/product_info"
	"client_streaming_pattern/server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", common.PortClientStreaming)
	if err != nil {
		log.Fatalf(common.LogPrefixClientStreaming+"failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, server.NewServer())
	log.Printf(common.LogPrefixClientStreaming+"Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
