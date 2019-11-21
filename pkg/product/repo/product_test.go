// Package repo
package repo

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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

func (s *UProductSuite) Test_Product_Purchase()  {
	if !testing.Short() {
		return
	}

	t := s.T()

	repo := &product{
		log: s.log.WithField("test_type", "unit_test"),
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


//endregion End test UProduct

func TestUProductSuite(t *testing.T) {


	suite.Run(t, &UProductSuite{
		log: logrus.WithField("pkg", "TestProductRepo"),
	})
}

//region Test IProduct
type IProductSuite struct {
	suite.Suite
	log *logrus.Entry
}

func (s *IProductSuite) SetupTest() {

}

func (s *IProductSuite) TearUpTest() {

}

func (s *IProductSuite) TearDownTest() {

}

func (s *IProductSuite) Test_Product_Purchase() {

	if testing.Short() {
		return
	}

	t := s.T()

	t.Run("quantity is not enough", func(t *testing.T) {

	})

	t.Run("row locked", func(t *testing.T) {

	})

	t.Run("should cause deadlock", func(t *testing.T) {

	})

	t.Run("success", func(t *testing.T) {

	})
}

//endregion End test IProduct

func TestIProductSuite(t *testing.T) {
	suite.Run(t, &IProductSuite{})
}

