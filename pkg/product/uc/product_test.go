// Package uc
package uc

import (
	"errors"
	"fmt"
	"github.com/ngdlong91/funtech/mocks"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"

	"github.com/ngdlong91/funtech/pkg/product/repo"

	"github.com/ngdlong91/funtech/dto"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

//region Test UProduct 
type UProductSuite struct {
	suite.Suite
	log *logrus.Entry
}

func (s *UProductSuite) SetupTest() {

}

func (s *UProductSuite) TearUpTest() {

}

func (s *UProductSuite) TearDownTest() {

}

func (s *UProductSuite) Test_Product_Purchase() {
	if !testing.Short() {
		return
	}

	t := s.T()
	c := &product{
		log: s.log,
	} // inject mock db and redis


	t.Run("multi request at same time and all success", func(t *testing.T) {
		repoMock := &mocks.ProductRepo{}
		repoMock.On("Purchase", 1, 1).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 2, 2).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 3, 3).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 1, 2).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 2, 3).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 3, 4).Return(dto.Product{}, nil)
		c.productRepo = repoMock

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
						Quantity: 2,
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
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
					{
						Id:       3,
						Quantity: 4,
					},
				},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also

		var wg sync.WaitGroup

		for _, req := range reqs {
			fmt.Printf("Start process for user %d \n", req.Id)
			wg.Add(1)
			go func() {
				res, err := c.Purchase(req)
				assert.Nil(t, err)
				assert.Equal(t, "success", res[0].Result)
				assert.Equal(t, "success", res[1].Result)
				assert.Equal(t, "success", res[2].Result)
				wg.Done()
			}()

		}
		wg.Wait()

	})

	t.Run("multi request and some success", func(t *testing.T) {
		repoMock := &mocks.ProductRepo{}
		repoMock.On("Purchase", 1, 1).Return(dto.Product{}, errors.New("out of stock"))
		repoMock.On("Purchase", 2, 2).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 3, 3).Return(dto.Product{}, errors.New("out of stock"))
		repoMock.On("Purchase", 1, 2).Return(dto.Product{}, nil)
		repoMock.On("Purchase", 2, 3).Return(dto.Product{}, errors.New("out of stock"))
		repoMock.On("Purchase", 3, 4).Return(dto.Product{}, nil)
		c.productRepo = repoMock

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
						Quantity: 2,
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
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
					{
						Id:       3,
						Quantity: 4,
					},
				},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also

		var wg sync.WaitGroup

		for _, req := range reqs {
			fmt.Printf("Start process for user %d \n", req.Id)
			wg.Add(1)
			go func() {
				res, err := c.Purchase(req)
				assert.Nil(t, err)
				if req.Id == 1 {
					assert.Equal(t, "out of stock", res[0].Result)
					assert.Equal(t, "success", res[1].Result)
					assert.Equal(t, "out of stock", res[2].Result)
				}

				if req.Id == 2 {
					assert.Equal(t, "success", res[0].Result)
					assert.Equal(t, "out of stock", res[1].Result)
					assert.Equal(t, "success", res[2].Result)
				}

				wg.Done()
			}()

		}
		wg.Wait()
	})
}

//endregion End test UProduct

func TestUProductSuite(t *testing.T) {
	suite.Run(t, &UProductSuite{
		log: logrus.WithField("pkg", "TestUseCases").WithField("type", "unit_test"),
	})
}

func TestProductUC_Purchase(t *testing.T) {

}

//region Integration test

func TestProductUC_IntegrationPurchase(t *testing.T) {

	if testing.Short() {
		return
	}

	productRepo := repo.NewProduct()
	c := product{
		log:         nil,
		productRepo: productRepo,
	}

	// Cause valid request should be done before call
	// so this test case will act as a note
	t.Run("should valid payload before called", func(t *testing.T) {
		err := errors.New("should be valid payload")
		assert.Error(t, errors.New("should be valid payload"), err)
	})

	t.Run("quantity not enough", func(t *testing.T) {
		payloadOneItem := dto.RequestPurchase{
			Id: 1,
			Products: []dto.Product{
				{
					Id:       1,
					Quantity: 9999999,
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
						Quantity: 2,
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
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
					{
						Id:       3,
						Quantity: 4,
					},
				},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also

		var wg sync.WaitGroup

		for _, req := range reqs {
			fmt.Printf("Start process for user %d \n", req.Id)
			wg.Add(1)
			go func() {
				res, err := c.Purchase(req)
				assert.Nil(t, err)
				assert.Equal(t, "success", res[0].Result)
				assert.Equal(t, "success", res[1].Result)
				assert.Equal(t, "success", res[2].Result)
				wg.Done()
			}()

		}
		wg.Wait()

	})

	t.Run("multi request and some success", func(t *testing.T) {
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
						Quantity: 2,
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
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
					{
						Id:       3,
						Quantity: 4,
					},
				},
			},
		}
		// Simulate multi request at the same time
		// Should make integration test and load test also

		var wg sync.WaitGroup

		for _, req := range reqs {
			fmt.Printf("Start process for user %d \n", req.Id)
			wg.Add(1)
			go func() {
				res, err := c.Purchase(req)
				assert.Nil(t, err)
				if req.Id == 1 {
					assert.Equal(t, "out of stock", res[0].Result)
					assert.Equal(t, "success", res[1].Result)
					assert.Equal(t, "out of stock", res[2].Result)
				}

				if req.Id == 2 {
					assert.Equal(t, "success", res[0].Result)
					assert.Equal(t, "out of stock", res[1].Result)
					assert.Equal(t, "success", res[2].Result)
				}

				wg.Done()
			}()

		}
		wg.Wait()
	})
}

//endregion Integration test
