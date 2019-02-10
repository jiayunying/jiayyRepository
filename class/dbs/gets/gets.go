package gets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"../conns" //引用数据库连接
	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(code, salary string) (string, error) {
	fmt.Println("正在执行查询Querymyself...")
	db, err := conns.OpenConn() //	打开数据库连接
	check(err)

	rows, err := db.Query("SELECT value,name FROM table1 where code='" + code + "' and salary='" + salary + "'")
	check(err)

	type myrow struct {
		value string
		name  string
	}

	//var tables []map[string]string
	// for rows.Next() {

	// 	var hello myrow
	// 	//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
	// 	if err := rows.Scan(&hello.value, &hello.name); err != nil {
	// 		log.Fatal("scan错误：", err)
	// 	}
	// 	//rows.Scan(&code, &value, &name, &salary),这样查询数据有错误，一定不能忘记用err接收rows.Scan的结果
	// 	//fmt.Printf(hello.code, hello.value, hello.name, hello.salary)
	// 	m := map[string]string{

	// 		"value": hello.value,
	// 		"name":  hello.name,
	// 	}
	// 	tables = append(tables)
	// }

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return "", err
	}

	user, err := RowsJson(rows)

	rows.Close()
	return user, err
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func RowsJson(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns() //取字段名
	if err != nil {
		return "", err
	}
	count := len(columns) //字段个数
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	jd := string(jsonData)
	fmt.Println(">>>Rowsjd<<<", jd)
	return jd, err
}
