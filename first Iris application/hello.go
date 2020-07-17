package hello

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"os"
	"strconv"
)


func hello() {
	app := iris.New()

	app.Get("/me", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("请求路径是：" + path)
	})
	
	app.Get("/userInfo", func(context *context.Context) {
		path := context.Path()
		app.Logger().Info(path)

		//获取请求参数
		userName := context.URLParam("userName")
		app.Logger().Info("用户姓名：" + userName)

		password := context.URLParam("password")
		app.Logger().Info("密码：" + password)

		passwordInt, err := strconv.Atoi(password)

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		if passwordInt == 1 {
			context.WriteString("登录成功")
		} else {
			context.WriteString("登录失败")
		}
	})
	
	app.Post("/list", func(c *context.Context) {
		path := c.Path()
		app.Logger().Info(path)

		name := c.PostValue("name")
		password := c.PostValue("password")

		c.WriteString("收到登录请求"+ name + password)
	})
	
	app.Post("/postJson", func(c *context.Context) {
		path := c.Path()
		app.Logger().Info("请求地址：", path)
		
		//json 解析
		var person Person
		if err := c.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		c.Writef("收到个人信息：%#+v\n", person)
	})

	app.Post("login", func(c *context.Context) {
		path := c.Path()
		app.Logger().Info("请求地址：", path)

		userId := c.PostValue("userId")
		password := c.PostValue("password")

		 result := false

		if userId == "Jack" && password == "123456" {
			result = true
		}

		c.JSON(iris.Map{"resquestCode" : "200","success":result})

	})

	app.Listen(":8000")
}

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type LoginInfo struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
}