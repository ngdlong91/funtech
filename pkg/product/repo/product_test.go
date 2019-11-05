// Package repo
package repo

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_product_Purchase(t *testing.T) {
	repo := &product{
		log: logrus.WithField("test", "productRepo"),
	}

	productColumns := []string{"id", "quantity"}

	t.Run("cannot begin transaction", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin().WillReturnError(errors.New("cannot begin error"))
		//mock.ExpectQuery("SELECT id, quantity from product where id = ? FOR SHARE").WithArgs(1).
		//	WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 1))
		//mock.ExpectExec("UPDATE product SET quantity = (quantity - ?) WHERE id = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		//mock.ExpectCommit()
		_, err = repo.Purchase(1, 777)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("db busy. Try later"), err)

	})

	t.Run("query not found/cannot find query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(-1).
			WillReturnRows(sqlmock.NewRows(productColumns))
		//mock.ExpectExec("UPDATE product SET quantity = (quantity - ?) WHERE id = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		//mock.ExpectCommit()
		_, err = repo.Purchase(-1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("invalid requests"), err)
	})

	t.Run("cannot update but rollback success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(1).
			WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 5))
		mock.ExpectExec(`UPDATE product SET quantity `).WithArgs(1, 1).WillReturnError(errors.New("cannot update"))
		mock.ExpectRollback()
		//mock.ExpectExec("UPDATE product SET quantity = (quantity - ?) WHERE id = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		//mock.ExpectCommit()
		_, err = repo.Purchase(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("db busy. Try later"), err)
	})

	t.Run("cannot update but rollback error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(1).
			WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 5))
		mock.ExpectExec(`UPDATE product SET quantity `).WithArgs(1, 1).WillReturnError(errors.New("cannot update"))
		mock.ExpectRollback().WillReturnError(errors.New("cannot rollback"))
		//mock.ExpectExec("UPDATE product SET quantity = (quantity - ?) WHERE id = ?").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		//mock.ExpectCommit()
		_, err = repo.Purchase(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("db busy. Try later"), err)
	})

	t.Run("update success but commit failed", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(1).
			WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 5))
		mock.ExpectExec(`UPDATE product SET quantity `).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit().WillReturnError(errors.New("cannot commit"))
		_, err = repo.Purchase(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("db busy. Try later"), err)
	})

	t.Run("commit failed and rollback failed", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(1).
			WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 5))
		mock.ExpectExec(`UPDATE product SET quantity `).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit().WillReturnError(errors.New("cannot commit"))
		mock.ExpectRollback().WillReturnError(errors.New("cannot rollback"))
		_, err = repo.Purchase(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("db busy. Try later"), err)
	})

	t.Run("out of stock", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(1).
			WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 5))
		_, err = repo.Purchase(1, 10)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("out of stock"), err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		repo.conn = db
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT id, quantity from product where id = ?`).WithArgs(1).
			WillReturnRows(sqlmock.NewRows(productColumns).AddRow(1, 5))
		mock.ExpectExec(`UPDATE product SET quantity `).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		_, err = repo.Purchase(1, 1)
		assert.Nil(t, err)
	})

}
