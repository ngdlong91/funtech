// Package uc
package uc

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ngdlong91/funtech/cmd/gin/dto"

	"github.com/stretchr/testify/assert"
)

func TestProductUC_Purchase(t *testing.T) {
	c := NewProductProcessor() // inject mock db and redis

	// Cause valid request should be done before call
	// so this test case will act as a note
	t.Run("should valid payload before called", func(t *testing.T) {
		err := errors.New("should be valid payload")
		assert.Error(t, errors.New("should be valid payload"), err)
	})

	t.Run("quantity not enough", func(t *testing.T) {
		payloadOneItem := dto.RequestPurchase{
			Id: 5,
			Products: []dto.Product{
				{
					Id:       1,
					Quantity: 5,
				},
			},
		}
		res, err := c.Purchase(payloadOneItem)
		assert.Error(t, errors.New("quantity not enough"), err)
		fmt.Printf("response %+v \n", res)

	})

	t.Run("multi request at same time and all success", func(t *testing.T) {
		reqs := []dto.RequestPurchase{
			{
				Id:       1,
				Products: []dto.Product{},
			},
			{
				Id:       2,
				Products: []dto.Product{},
			},
			{
				Id:       3,
				Products: []dto.Product{},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also
		for _, req := range reqs {
			go func() {
				responses, err := c.Purchase(req)
				assert.Error(t, nil, err)
				for _, res := range responses {
					assert.Equal(t, dto.PurchaseResult{Id: req.Id, IsSuccess: true}, res)
				}
			}()
		}
	})

	t.Run("multi request and some success", func(t *testing.T) {
		reqs := []dto.RequestPurchase{
			{
				Id:       1,
				Products: []dto.Product{},
			},
			{
				Id:       2,
				Products: []dto.Product{},
			},
			{
				Id:       3,
				Products: []dto.Product{},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also
		for _, req := range reqs {
			go func() {
				responses, err := c.Purchase(req)
				assert.Error(t, nil, err)
				for _, res := range responses {
					assert.Equal(t, dto.PurchaseResult{Id: req.Id, IsSuccess: true}, res)
				}
			}()
		}
	})

	t.Run("multi request and all failed", func(t *testing.T) {
		reqs := []dto.RequestPurchase{
			{
				Id:       1,
				Products: []dto.Product{},
			},
			{
				Id:       2,
				Products: []dto.Product{},
			},
			{
				Id:       3,
				Products: []dto.Product{},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also
		for _, req := range reqs {
			go func() {
				responses, err := c.Purchase(req)
				assert.Error(t, nil, err)
				for _, res := range responses {
					assert.Equal(t, dto.PurchaseResult{Id: req.Id, IsSuccess: true}, res)
				}
			}()
		}
	})
}
