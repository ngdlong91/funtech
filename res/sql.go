// Package res
package res

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SQL struct {
	db *sql.DB
}

func connect() *sql.DB {
	fmt.Println("make connect to mysql server")
	db, err := sql.Open("mysql", getInfo())
	if err != nil {
		fmt.Println("cannot connect to sql")
		return nil
	}

	fmt.Println("Open sql connection")

	rows, err := db.Query("SELECT quantity FROM product WHERE id = 1")
	if err != nil {
		fmt.Println("Cannot query ", err.Error())
		panic(err)
	}

	fmt.Println("Query with select")

	for rows.Next() {
		var quantity int
		if err := rows.Scan(&quantity); err != nil {
			panic(err)
		}

		fmt.Printf("Obj %d \n", quantity)

	}
	return db

}

func getInfo() string {
	return "root:123456@tcp(localhost:32769)/funtech"
}

func (s *SQL) Conn() *sql.DB {
	return s.db
}

func NewSQLInstance() *SQL {
	return &SQL{
		db: connect(),
	}
}
