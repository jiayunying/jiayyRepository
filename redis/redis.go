package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//redis连接
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	//redis读写，永久写入
	_, err = c.Do("SET", "mykey", "jiayy")
	if err != nil {
		fmt.Println("redis set faild", err)
		return
	}
	//判断某键是否存在
	isExisit, err := redis.Bool(c.Do("EXISTS", "mykey"))
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("exists or not:%v \n", isExisit)

	//读取某键对应的值,c.Do(不区分大小写)
	mykey, err := redis.String(c.Do("Get", "mykey"))
	if err != nil {
		fmt.Println("redis get faild ", err)
		return
	}
	fmt.Printf("Get mykey:%v \n", mykey)

	//删除某键值对
	// _,err:=c.Do("DEL","mykey")
	// if err!=nil{
	// 	fmt.Println("del faild",err)
	// }

	//往redis中读写json
	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)
	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}
	var imapGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}
