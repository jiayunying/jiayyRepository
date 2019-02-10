package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome!</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		fmt.Println("正在执行查询query...")
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
		check(err)

		rows, err := db.Query("SELECT * FROM test.table1")
		check(err)
		record := make(map[string]string) //返回多个数据集
		for rows.Next() {
			columns, _ := rows.Columns()

			scanArgs := make([]interface{}, len(columns)) //make创建切片
			values := make([]interface{}, len(columns))

			for i := range values {
				scanArgs[i] = &values[i]
			}

			//将数据保存到 record 字典
			err = rows.Scan(scanArgs...)

			for i, col := range values {
				if col != nil {
					record[columns[i]] = string(col.([]byte))
				}
			}
			fmt.Println(record)
		}
		rows.Close()

		ctx.JSON(record)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
