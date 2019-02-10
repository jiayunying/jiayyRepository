package conns

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//var TESTDB *DB

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	return db, err
}
