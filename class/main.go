package main

import (
	"./rests"

	"github.com/kataras/iris"
)

func main() {
	//fmt.Println("================")
	//gets.GetAllUsers()
	//fmt.Println("<<<<<<<<<<<<<<<<<<<")
	//gets.Insert("1111","MMMM")
	//fmt.Println("+++++++++++++++++++")
	//gets.GetAllUsers()
	app := iris.New()
	rests.UserRest(app)
	app.Run(iris.Addr(":8080"))
}
