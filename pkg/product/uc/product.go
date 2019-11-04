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
		product, err := c.productRepo.Select(item.Id)
		if err != nil {
			result = dto.PurchaseResult{
				Id:        item.Id,
				IsSuccess: false,
			}
			results = append(results, result)
			continue
		}
		c.log.Infof("Product details %+v \n", product)

		if product.Quantity < item.Quantity {
			c.log.Errorf("product quantity is not enough")
			result = dto.PurchaseResult{
				Id:        item.Id,
				IsSuccess: false,
			}
			results = append(results, result)
			continue
		}

		if _, err := c.productRepo.Update(item.Id, item.Quantity); err != nil {
			c.log.Errorf("Update product err: %s \n", err.Error())
			result = dto.PurchaseResult{
				Id:        item.Id,
				IsSuccess: false,
			}
			results = append(results, result)
			continue
		}

		result = dto.PurchaseResult{
			Id:        item.Id,
			IsSuccess: true,
		}

		c.log.Infof("Purchased item %+v \n", item)

		results = append(results, result)
	}

	c.log.Infof("Final result %+v \n", results)

	return results, nil

}

func NewProductProcessor() Product {
	return &product{
		log:         logrus.WithField("uc", "product"),
		productRepo: repo.NewProduct(),
	}
}
