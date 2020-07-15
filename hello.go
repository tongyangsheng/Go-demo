package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Get("/", index)

	app.Get("/me", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("请求路径是：" + path)
	})

	app.Listen(":8000")
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>Hello, World!</h1>")
}