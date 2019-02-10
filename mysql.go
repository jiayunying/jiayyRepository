package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//insert()      //插入
	//query2() //查询
	//querymyself() //查询
	update() //修改
	//remove()  //删除

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func query() {
	fmt.Println("正在执行查询query...")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	check(err)

	rows, err := db.Query("SELECT * FROM test.table1")
	check(err)

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns)) //make创建切片
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()
}
func querymyself() {
	fmt.Println("正在执行查询Querymyself...")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	check(err)

	rows, err := db.Query("SELECT code,value,name,salary FROM table1 where code='1001'")
	check(err)

	for rows.Next() {

		type myrow struct {
			code   string
			value  sql.NullString
			name   string
			salary string
		}

		var hello myrow
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&hello.code, &hello.value, &hello.name, &hello.salary); err != nil {
			log.Fatal("scan错误：", err)
		}
		//rows.Scan(&code, &value, &name, &salary),这样查询数据有错误，一定不能忘记用err接收rows.Scan的结果
		fmt.Printf(hello.code, hello.value.String, hello.name, hello.salary)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}
func query2() {
	fmt.Println("正在执行查询Query2...")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	check(err)

	rows, err := db.Query("SELECT code,value,name,salary FROM table1 where ")
	check(err)

	for rows.Next() {

		var code string
		var value sql.NullString //如何处理空值？
		var name string
		var salary string
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&code, &value.String, &name, &salary); err != nil {
			log.Fatal("scan错误：", err)
		}
		if value.Valid {
			fmt.Printf("code: %s value: %s name: %s salary: %s\n", code, value.String, name, salary)
		} else {
			fmt.Printf("code: %s value: '' name: %s salary: %s\n", code, name, salary)
		}
		//rows.Scan(&code, &value, &name, &salary),这样查询数据有错误，一定不能忘记用err接收rows.Scan的结果
		//fmt.Printf("code: %s value: %s name: %s salary: %s\n", code, value.String, name, salary)
		//fmt.Printf("code:", code, "value:", value.Valid?value.String:nil, "name:", name, "salary:", salary)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}

func insert() {
	fmt.Println("正在执行插入操作insert")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	check(err)

	stmt, err := db.Prepare(`INSERT table1 (code, value, name, salary) VALUES (?, ?, ?, ?)`)
	check(err)

	res, err := stmt.Exec("1006", nil, "李六", 100)
	check(err)

	id, err := res.LastInsertId()
	check(err)

	fmt.Println("插入数据成功：", id)
	stmt.Close()
}

func update() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	check(err)

	stmt, err := db.Prepare("UPDATE table1 set value=?, name=? WHERE code=?")
	check(err)

	res, err := stmt.Exec("修改", "修改", 1004)
	check(err)

	num, err := res.RowsAffected()
	check(err)

	fmt.Println("修改了 ", num, "行")
	stmt.Close()
}

func remove() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	check(err)

	stmt, err := db.Prepare("DELETE FROM table1 WHERE id=?")
	check(err)

	res, err := stmt.Exec(1004)
	check(err)

	num, err := res.RowsAffected()
	check(err)

	fmt.Println(num)
	stmt.Close()
}
