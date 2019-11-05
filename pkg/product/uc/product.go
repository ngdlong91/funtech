package uc

import (
	"github.com/sirupsen/logrus"

	"github.com/ngdlong91/funtech/dto"
	"github.com/ngdlong91/funtech/pkg/product/repo"
)

type ProductUseCase interface {
	Purchase(payload dto.RequestPurchase) ([]dto.PurchaseResult, error)
}

type product struct {
	log         *logrus.Entry
	productRepo repo.ProductRepo
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
			Result:    "success",
		}

		c.log.Debugf("Purchased item %+v \n", item)

		results = append(results, result)
	}

	c.log.Debugf("Final result %+v \n", results)

	return results, nil

}

func NewProductProcessor() ProductUseCase {
	return &product{
		log:         logrus.WithField("uc", "product"),
		productRepo: repo.NewProduct(),
	}
}
