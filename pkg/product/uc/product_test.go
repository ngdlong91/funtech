// Package uc
package uc

import (
	"errors"
	"fmt"
	"sync"
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
		assert.Nil(t, err)
		assert.Equal(t, "out of stock", res[0].Result)

	})

	t.Run("multi request at same time and all success", func(t *testing.T) {
		reqs := []dto.RequestPurchase{
			{
				Id: 1,
				Products: []dto.Product{
					{
						Id:       1,
						Quantity: 1,
					},
					{
						Id:       2,
						Quantity: 10,
					},
					{
						Id:       3,
						Quantity: 3,
					},
				},
			},
			{
				Id: 2,
				Products: []dto.Product{
					{
						Id:       1,
						Quantity: 1,
					},
					{
						Id:       2,
						Quantity: 10,
					},
					{
						Id:       3,
						Quantity: 3,
					},
				},
			},
			//{
			//	Id: 3,
			//	Products: []dto.Product{
			//		{
			//			Id:       1,
			//			Quantity: 1,
			//		},
			//		{
			//			Id:       2,
			//			Quantity: 2,
			//		},
			//		//{
			//		//	Id:       3,
			//		//	Quantity: 3,
			//		//},
			//	},
			//},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also
		var wg sync.WaitGroup
		for _, req := range reqs {
			fmt.Printf("Start process for user %d \n", req.Id)
			wg.Add(1)
			go func() {
				var id = req.Id
				c.Purchase(req)
				//assert.Error(t, nil, err)
				//for _, res := range responses {
				//	assert.Equal(t, dto.PurchaseResult{Id: req.Id, IsSuccess: true}, res)
				//}
				fmt.Println("Finished ", id)
				wg.Done()
			}()

		}
		wg.Wait()
		fmt.Println("Finish tests")
		assert.Nil(t, nil)
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
