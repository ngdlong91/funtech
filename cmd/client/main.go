package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ngdlong91/funtech/cmd/gin/pkg/product"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := product.NewProductClient(conn)

	req, err := buildRequestFromArgs(os.Args)
	if err != nil {
		panic(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DoPurchase(ctx, &req)
	if err != nil {
		log.Fatalf("could not purchases: %v", err)
	}
	log.Printf("Purchase result: %+v \n", r.Response)
}

func buildRequestFromArgs(args []string) (product.PurchaseRequest, error) {
	var req product.PurchaseRequest

	fmt.Printf("Input params %+v \n", args)
	// Make sure the
	size := len(args)
	if size > 1 && (size-2)%2 == 0 {
		cusId, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return product.PurchaseRequest{}, err
		}
		req.CustomerId = int32(cusId)
		for i := 2; i < size; i = i + 2 {
			productId, err := strconv.ParseInt(args[i], 10, 32)
			if err != nil {
				return product.PurchaseRequest{}, err
			}
			purchaseQuantity, err := strconv.ParseInt(args[i+1], 10, 32)
			if err != nil {
				return product.PurchaseRequest{}, err
			}
			productDetail := product.ProductDetail{
				Id:       int32(productId),
				Quantity: int32(purchaseQuantity),
			}

			req.Products = append(req.Products, &productDetail)
		}
		return req, nil

	}

	return product.PurchaseRequest{
		CustomerId: 1,
		Products: []*product.ProductDetail{
			{
				Id:       1,
				Quantity: 1,
			},
		},
	}, nil
}
