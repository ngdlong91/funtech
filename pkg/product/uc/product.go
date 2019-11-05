package uc

import (
	"github.com/sirupsen/logrus"

	"github.com/ngdlong91/funtech/cmd/gin/pkg/product/repo"

	"github.com/ngdlong91/funtech/cmd/gin/dto"
)

type Product interface {
	Purchase(payload dto.RequestPurchase) ([]dto.PurchaseResult, error)
}

type product struct {
	log          *logrus.Entry
	productRepo  repo.Product
	productCache repo.ProductCache
}

// Purchase process request from client
// payload should be pre-process and validate before call this function
// so we skip validate
func (c *product) Purchase(payload dto.RequestPurchase) ([]dto.PurchaseResult, error) {
	var results []dto.PurchaseResult
	for _, item := range payload.Products {
		var result dto.PurchaseResult
		c.log.Debugf("Start handle purchase request...")
		if _, err := c.productRepo.Purchase(item.Id, item.Quantity); err != nil {
			c.log.Errorf("Purchase product err: %s \n", err.Error())
			result = dto.PurchaseResult{
				Id:        item.Id,
				IsSuccess: false,
				Result:    "out of stock", //Out of stock
			}
			results = append(results, result)
			continue
		}

		result = dto.PurchaseResult{
			Id:        item.Id,
			IsSuccess: true,
			Result:    "purchased",
		}

		c.log.Debugf("Purchased item %+v \n", item)

		results = append(results, result)
	}

	c.log.Debugf("Final result %+v \n", results)

	return results, nil

}

func NewProductProcessor() Product {
	return &product{
		log:         logrus.WithField("uc", "product"),
		productRepo: repo.NewProduct(),
	}
}
