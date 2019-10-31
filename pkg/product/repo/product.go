package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ngdlong91/funtech/cmd/gin/dto"
)

type Product interface {
	Select(id int) (dto.Product, error)
	Update(id, quantity int) (dto.Product, error)
}

// cache is a temp memory storage.
type cache struct {
	// redis
	isDown bool
}

func (r *cache) Select(id int) (dto.Product, error) {
	// Conn fail for some reason
	if r.isDown {
		return dto.Product{}, errors.New("product cache has problems")
	}

	return dto.Product{
		Id:       1,
		Quantity: 1,
	}, nil
}

func (r *cache) Update(id, quantity int) (dto.Product, error) {
	return dto.Product{}, nil
}

type product struct {
	tempRepo cache
	isDown   bool

	conn sql.Conn

	ctx context.Context
}

func (r *product) Select(id int) (dto.Product, error) {
	product, err := r.tempRepo.Select(id)
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

		if _, err := r.tempRepo.Update(id, product.Quantity); err != nil {
			// todo: handle case cannot update temp repo
		}
		return product, nil
	}

	return product, nil
}

func (r *product) doDBSelect(id int) (dto.Product, error) {
	return dto.Product{}, nil
}

func (r *product) Update(id, quantity int) (dto.Product, error) {
	if _, err := r.tempRepo.Update(id, quantity); err != nil {
		// Todo: Handle case temp storage got problem. Go back or turn flag for this
	}

	// Do db update
	tx, err := r.conn.BeginTx(r.ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		return dto.Product{}, err
	}

	if _, err := tx.Exec("SELECT quantity from product where id = ? for update ", id); err != nil {
		if err := tx.Rollback(); err != nil {
			return dto.Product{}, err
		}
		return dto.Product{}, err
	}

	if _, err := tx.Exec(`UPDATE product SET quantity = ? WHERE id = ?`, quantity, id); err != nil {
		if err := tx.Rollback(); err != nil {
			return dto.Product{}, err
		}
		return dto.Product{}, err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return dto.Product{}, err
		}
		return dto.Product{}, err
	}

	// Return new product data from db
	return r.Select(id)
}

func NewProduct() Product {
	return nil
}
