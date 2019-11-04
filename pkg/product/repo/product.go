package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ngdlong91/funtech/cmd/gin/res"

	"github.com/ngdlong91/funtech/cmd/gin/dto"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

var productKey = "product#%d"

type Product interface {
	Insert(quantity int) error
	Select(id int) (dto.Product, error)
	Update(id, quantity int) (dto.Product, error)
}

// cache is a temp memory storage.
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

	fmt.Printf("Query result %+v \n", result)
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

func (r *product) Update(id, quantity int) (dto.Product, error) {
	if _, err := r.cache.Update(id, quantity); err != nil {

		r.log.Errorf("cannot update cache err: %s \n", err.Error())
		// Todo: Handle case temp storage got problem. Go back or turn flag for this
	}

	// Do db update
	tx, err := r.conn.Begin()
	if err != nil {
		r.log.Errorf("begin transaction error \n", err.Error())
		return dto.Product{}, err
	}

	if _, err := tx.Exec("SELECT quantity from product where id = ? for update ", id); err != nil {
		r.log.Errorf("cannot lock cols err: %s \n", err.Error())
		if err := tx.Rollback(); err != nil {
			r.log.Errorf("Rollback err: %s \n", err.Error())
			return dto.Product{}, err
		}
		return dto.Product{}, err
	}

	if _, err := tx.Exec(`UPDATE product SET quantity = (quantity - ?) WHERE id = ?`, quantity, id); err != nil {
		r.log.Errorf("cannot update record err: %s \n", err.Error())
		if err := tx.Rollback(); err != nil {
			r.log.Errorf("rollback err: %s \n", err.Error())
			return dto.Product{}, err
		}
		return dto.Product{}, err
	}

	if err := tx.Commit(); err != nil {
		r.log.Errorf("cannot commit err: %s \n", err.Error())
		if err := tx.Rollback(); err != nil {
			r.log.Errorf("rollback err: %s \n", err.Error())
			return dto.Product{}, err
		}
		return dto.Product{}, err
	}

	// Return new product data from db
	return r.Select(id)
}

func NewProduct() Product {
	return &product{
		cache: newProductCache(),
		log:   logrus.WithField("services", "product"),
		conn:  res.NewSQLInstance().Conn(),
	}
}
