package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/ngdlong91/funtech/pkg/product"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var serverAddress string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		return
	}
	serverAddress = os.Getenv("GRPC_SERVER")

	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	fmt.Println("Start grpc server")
	lis, err := net.Listen("tcp", serverAddress)
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
