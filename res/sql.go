// Package res
package res

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SQL struct {
	db *sql.DB
}

func connect() *sql.DB {

	db, err := sql.Open("mysql", getInfo())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to mysql server")
	return db

}

func getInfo() string {

	server := os.Getenv("SQL_SERVER")
	if server == "" {
		server = "root:123456@tcp(localhost:32769)/funtech"

	}
	return server
}

func (s *SQL) Conn() *sql.DB {
	return s.db
}

func NewSQLInstance() *SQL {
	return &SQL{
		db: connect(),
	}
}
