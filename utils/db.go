package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/cloud")
	if err != nil {
		panic(err.Error())
	}
}
