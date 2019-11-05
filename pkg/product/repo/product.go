package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ngdlong91/funtech/dto"
	"github.com/ngdlong91/funtech/res"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

var productKey = "product#%d"

type ProductRepo interface {
	Insert(quantity int) error
	Select(id int) (dto.Product, error)
	Purchase(id, quantity int) (dto.Product, error)
}

// cache is a temp memory storage.s
type cache struct {
	// redis
	redis res.CRedis
}

type product struct {
	log *logrus.Entry

	cache  ProductCache
	isDown bool

	conn *sql.DB

	ctx context.Context
}

func (r *product) Insert(quantity int) error {
	query := `INSERT INTO product (quantity, created_at) VALUES (?, ?)`
	if _, err := r.conn.Exec(query, quantity, time.Now().Unix()); err != nil {
		return err
	}
	return nil
}

func (r *product) Select(id int) (dto.Product, error) {
	return r.doDBSelect(id)
	product, err := r.cache.Select(id)
	if err != nil {
		// Get from database
		if r.isDown {
			return dto.Product{}, errors.New("product db has problems")
		}
		return r.doDBSelect(id)
	}

	// Temp repo work but we still select from not init key
	if product.Id == 0 {
		product, err := r.doDBSelect(id)
		if err != nil {
			return dto.Product{}, err
		}

		if _, err := r.cache.Update(id, product.Quantity); err != nil {
			// todo: handle case cannot update temp repo
		}
		return product, nil
	}

	return product, nil
}

func (r *product) doDBSelect(id int) (dto.Product, error) {
	result, err := r.conn.Query("SELECT id, quantity FROM product WHERE id = ?", id)
	if err != nil {
		r.log.Errorf("cannot select product err: %s \n", err.Error())
		return dto.Product{}, err
	}

	for result.Next() {
		var product dto.Product
		if err := result.Scan(&product.Id, &product.Quantity); err != nil {
			r.log.Errorf("Scan product error: %s \n", err.Error())
			return dto.Product{}, err
		}

		return product, nil
	}

	return dto.Product{}, errors.New("record not found")
}

func (r *product) Purchase(id, quantity int) (dto.Product, error) {
	//if _, err := r.cache.Update(id, quantity); err != nil {
	//	r.log.Errorf("cannot update cache err: %s \n", err.Error())
	//	// Todo: Handle case temp storage got problem. Go back or turn flag for this
	//}

	r.log.Debugf("begin transaction, try to purchase id %d with quantity %d \n", id, quantity)
	// Do db update
	tx, err := r.conn.Begin()
	if err != nil {
		r.log.Errorf("begin transaction error: %s \n", err.Error())
		return dto.Product{}, errors.New("db busy. Try later")
	}
	r.log.Debugf("Try to get details")
	rows, err := tx.Query("SELECT id, quantity from product where id = ? FOR SHARE", id)
	if err != nil {
		r.log.Errorf("Get product details %s \n", err.Error())
		return dto.Product{}, errors.New("cannot get product detail ")
	}

	var product dto.Product
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Quantity); err != nil {
			return dto.Product{}, err
		}
	}
	if err := rows.Close(); err != nil {
		return dto.Product{}, err
	}

	if product.Id == 0 {
		return dto.Product{}, errors.New("invalid requests")
	}
	r.log.Debugf("ProductRepo details %+v \n", product)

	if product.Quantity >= quantity {
		if _, err := tx.Exec(`UPDATE product SET quantity = (quantity - ?) WHERE id = ?`, quantity, id); err != nil {
			r.log.Errorf("cannot update record err: %s \n", err.Error())
			if err := tx.Rollback(); err != nil {
				r.log.Errorf("rollback err: %s \n", err.Error())
				return dto.Product{}, errors.New("db busy. Try later")
			}
			return dto.Product{}, errors.New("db busy. Try later")
		}

		if err := tx.Commit(); err != nil {
			r.log.Errorf("cannot commit err: %s \n", err.Error())
			if err := tx.Rollback(); err != nil {
				r.log.Errorf("rollback err: %s \n", err.Error())
				return dto.Product{}, errors.New("db busy. Try later")
			}
			return dto.Product{}, errors.New("db busy. Try later")
		}

		return dto.Product{}, nil
	}
	return dto.Product{}, errors.New("out of stock")
}

func NewProduct() ProductRepo {
	return &product{
		cache: newProductCache(),
		log:   logrus.WithField("services", "product"),
		conn:  res.NewSQLInstance().Conn(),
	}
}
