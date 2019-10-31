// Package uc
package uc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductUC_Purchase(t *testing.T) {
	c := product{} // inject mock db and redis
	t.Run("invalid request", func(t *testing.T) {

		c.Purchase()
	})

	t.Run("quantity not enough", func(t *testing.T) {
		reqQuantity := 10
		assert.Equal(t, 10, reqQuantity)
	})

	t.Run("multi request at same time and all success", func(t *testing.T) {

	})

	t.Run("multi request and some success", func(t *testing.T) {

	})

	t.Run("multi request and all failed", func(t *testing.T) {

	})
}
