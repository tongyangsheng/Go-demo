package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()

	/*
	handle
	 */

	//GET
	app.Handle("GET","/userInfo", func(c *context.Context) {
		path := c.Path()
		app.Logger().Info("connect url:",path)
		app.Logger().Error("error:","request time out")
	})

	//POST
	app.Handle("POST", "/postCommit", func(c *context.Context) {
		path := c.Path()
		app.Logger().Info("connect url:",path)

	})

	/*
	正则表达
	 */

	app.Get("/city/{cityName}", func(c *context.Context) {
		path := c.Path()
		app.Logger().Info("connect url:",path)

		cityName := c.Params().Get("cityName")
		c.Writef("city is: %v",cityName)
	})

	app.Get("/loginName/{userId:int}", func(c *context.Context) {

		userName, err := c.Params().GetInt("userId")

		if err != nil {
			c.JSON(map[string]interface{}{
				"requestCode": 201,
				"message":"user类型错误",
			})
			return
		}

		c.JSON(map[string]interface{}{
		    "requestCode": 201,
			"message": userName,
		})
	})

	app.Listen(":8001")
}