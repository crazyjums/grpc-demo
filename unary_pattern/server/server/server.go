package server

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"unary_pattern/common"
	pb "unary_pattern/pb/genproto/product_info"
)

type Server struct {
	pb.UnimplementedProductInfoServer
	productMap map[string]*pb.Product
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) mustEmbedUnimplementedOrderManagementServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetProduct(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	product, ok := s.productMap[in.Value]
	if !ok {
		return nil, errors.New("product is not exist! ")
	}
	log.Printf(common.LogPrefixUnary+"GetProduct is matching product %+v\n", product)

	return product, nil
}

func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductId, error) {
	productId, err := generateProductId()
	if err != nil {
		log.Fatalf("failed to generate product id: %v", err)
	}

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	product := &pb.Product{
		Id:    productId,
		Name:  in.Name,
		Price: in.Price,
		Desc:  in.Desc,
	}
	s.productMap[productId] = product
	log.Printf(common.LogPrefixUnary+"AddProduct %+v\n", product)
	log.Printf("ProductMap %+v\n", s.productMap)

	return &pb.ProductId{
		Value: productId,
	}, nil
}

func generateProductId() (string, error) {
	// 生成一个版本为4的UUID
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // 设置版本为4
	uuid[8] = (uuid[8] & 0xbf) | 0x80 // 设置变体标识符

	// 将UUID转换为12位数字长度的id
	id := ""
	for i := 0; i < 6; i++ {
		value := uint64(uuid[i*2])<<8 + uint64(uuid[i*2+1])
		id += fmt.Sprintf("%05d", value%100000)
	}

	return id, nil
}
