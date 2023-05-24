package server

import (
	pb "client_streaming_pattern/pb/genproto/product_info"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
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
	log.Printf("GetProduct %s\n", in.Value)
	product, ok := s.productMap[in.Value]
	if !ok {
		return nil, errors.New("product is not exist! ")
	}

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
	log.Printf("AddProduct %+v\n", product)

	return &pb.ProductId{
		Value: productId,
	}, nil
}

func (s *Server) SearchProducts(in *pb.ProductName, stream pb.ProductInfo_SearchProductsServer) error {
	if in == nil || in.Value == "" {
		return errors.New("invalid product name")
	}

	log.Printf("SearchProducts %s, curr productMap: %+v \n", in.Value, s.productMap)
	for _, product := range s.productMap {
		if strings.Contains(product.Name, in.Value) {
			err := stream.Send(product)
			if err != nil {
				log.Fatalf("failed to send: %v", err)
				return err
			}
			log.Printf("match product found key= %+v\n", in.Value)
			time.Sleep(time.Millisecond * 500)
		}
	}

	return nil
}

func (s *Server) UpdateProducts(stream pb.ProductInfo_UpdateProductsServer) error {
	updatedStr := "updated..."
	for {
		product, err := stream.Recv()
		log.Printf("UpdateProducts %+v\n", product)

		if err == io.EOF || product == nil {
			return stream.SendAndClose(&pb.ProductId{
				Value: "",
			})
		}

		product.Name = product.Name + updatedStr
		s.productMap[product.Id] = product
		log.Printf("UpdateProducts %+v\n", product)

		if err != nil {
			return err
		}
	}
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
