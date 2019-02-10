package gets

import (
	"fmt"

	"../conns" //引用数据库连接
	_ "github.com/go-sql-driver/mysql"
)

func Insert(code string, value string) {
	fmt.Println("正在执行插入操作insert")

	db, err := conns.OpenConn() //	打开数据库连接
	check(err)

	stmt, err := db.Prepare(`INSERT table1 (code, value) VALUES (?, ?)`)
	check(err)

	_, err = stmt.Exec(code, value)
	check(err)

	stmt.Close()
}
