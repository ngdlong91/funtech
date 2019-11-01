// Package uc
package uc

import (
	"errors"
	"testing"

	"github.com/ngdlong91/funtech/cmd/gin/dto"

	"github.com/stretchr/testify/assert"
)

func TestProductUC_Purchase(t *testing.T) {
	c := product{} // inject mock db and redis
	payload := dto.RequestPurchase{
		Id: 0,
		Products: []dto.Product{
			{
				Id:       1,
				Quantity: 5,
			},
		},
	}
	t.Run("invalid request", func(t *testing.T) {
		err := c.Purchase(payload)
		assert.Error(t, errors.New("request invalid"), err)
	})

	t.Run("quantity not enough", func(t *testing.T) {
		err := c.Purchase(payload)
		assert.Error(t, errors.New("quantity not enough"), err)
	})

	t.Run("multi request at same time and all success", func(t *testing.T) {

	})

	t.Run("multi request and some success", func(t *testing.T) {

	})

	t.Run("multi request and all failed", func(t *testing.T) {

	})
}
