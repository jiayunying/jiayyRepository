package rests

import (
	"../dbs/gets"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func UserRest(app *iris.Application) {
	//app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Post("/home", func(ctx iris.Context) {

			type User struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}

			c := &User{}

			if err := ctx.ReadJSON(c); err != nil {
				panic(err.Error())
			} else {
				//ctx.JSON(c);
				json, _ := gets.GetUsers(c.Username, c.Password)
				ctx.WriteString(json)
			}

		})
		v1.Get("/about", func(ctx iris.Context) {
			ctx.WriteString("Hello from /about")
		})
		v1.Post("/send", func(ctx iris.Context) {
			ctx.WriteString("sent")
		})
		v1.Put("/send", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/send", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
	}

	app.Run(iris.Addr(":8080"))
}
