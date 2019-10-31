// Package repo
package repo

import (
	"errors"
	"testing"

	"github.com/ngdlong91/funtech/cmd/gin/dto"
	"github.com/stretchr/testify/assert"
)

func Test_temp_Select(t *testing.T) {
	repo := &product{}
	t.Run("temp repo error", func(t *testing.T) {
		product, err := repo.Select(777)
		assert.Nil(t, err)
		assert.Equal(t, dto.Product{
			Id:       0,
			Quantity: 0,
		}, product)

	})

	t.Run("temp repo not setup", func(t *testing.T) {

	})

	t.Run("temp repo response successs", func(t *testing.T) {

	})
}

func Test_product_Select(t *testing.T) {
	repo := product{}
	t.Run("temp repo error", func(t *testing.T) {
		product, err := repo.Select(5)
		assert.Error(t, errors.New("cannot get data from temp storage"), err)
		assert.Equal(t, dto.Product{}, product)
	})

	t.Run("temp repo is not setup", func(t *testing.T) {
		err := errors.New("not implement")

		assert.Equal(t, errors.New("temp repo is not setup"), err)

	})

	t.Run("temp repo response success", func(t *testing.T) {
		err := errors.New("not implement")
		assert.Equal(t, errors.New("temp repo response successs"), err)
	})
}

func Test_product_Update(t *testing.T) {

}
