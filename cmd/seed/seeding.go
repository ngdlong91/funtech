package main

import (
	"math/rand"

	"github.com/ngdlong91/funtech/pkg/product/repo"
)

const NumberOfRecord int = 200

func main() {
	productRepo := repo.NewProduct()
	for i := 0; i < NumberOfRecord; i++ {
		quantity := rand.Intn(10) + 1
		if err := productRepo.Insert(quantity); err != nil {

		}
	}
}
