package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"unary_pattern/common"
	pb "unary_pattern/pb/genproto/product_info"
	"unary_pattern/server/server"
)

func main() {
	lis, err := net.Listen("tcp", common.PortUnary)
	if err != nil {
		log.Fatalf(common.LogPrefixUnary+"failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, server.NewServer())
	log.Printf(common.LogPrefixUnary+"Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf(common.LogPrefixUnary+"failed to serve: %v", err)
	}
}
