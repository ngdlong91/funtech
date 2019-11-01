package uc

import (
	"fmt"

	"github.com/ngdlong91/funtech/cmd/gin/pkg/product/repo"

	"github.com/ngdlong91/funtech/cmd/gin/dto"
)

type Product interface {
	Purchase(payload dto.RequestPurchase) error
}

type product struct {
	productRepo repo.Product
}

func (c *product) Purchase(payload dto.RequestPurchase) error {
	for _, item := range payload.Products {
		fmt.Printf("Try to purchase item %+v \n", item)
	}

	return nil

}

func NewProductProcessor() Product {
	return &product{}
}
