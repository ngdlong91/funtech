package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ngdlong91/funtech/cmd/gin/pkg/product"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start grpc server")
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	productHandler := product.NewRPCHandler()
	product.RegisterProductServer(s, productHandler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
