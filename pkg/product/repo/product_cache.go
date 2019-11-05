// Package repo
package repo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ngdlong91/funtech/cmd/gin/dto"

	"github.com/ngdlong91/funtech/cmd/gin/res"
)

type ProductCache interface {
	Select(id int) (dto.Product, error)
	Update(id, quantity int) (dto.Product, error)
}

type productCache struct {
	log   *logrus.Entry
	redis res.CRedis
}

func (r *productCache) Select(id int) (dto.Product, error) {
	// Conn fail for some reason
	if r.redis.IsRun() {
		return dto.Product{}, errors.New("product cache has problems")
	}

	valBytes, err := r.redis.Get(fmt.Sprintf(productKey, id))
	if err != nil {
		return dto.Product{}, err
	}

	var product dto.Product
	if err := json.Unmarshal(valBytes, &product); err != nil {
		return dto.Product{}, err
	}

	fmt.Printf("Result from cache %+v \n", product)
	return product, nil
}

func (r *productCache) Update(id, quantity int) (dto.Product, error) {
	key := fmt.Sprintf(productKey, id)

	product := dto.Product{Id: id, Quantity: quantity}

	valStr, err := json.Marshal(&product)
	if err != nil {
		return dto.Product{}, err
	}

	err = r.redis.Set(key, valStr)
	if err != nil {
		return dto.Product{}, err
	}

	return dto.Product{}, nil
}

func newProductCache() ProductCache {
	return &productCache{
		log:   logrus.WithField("cache", "product"),
		redis: res.RedisInstance(),
	}
}
