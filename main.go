package main

import (
	"github.com/kataras/iris"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
	"./controller"
	"github.com/kataras/iris/sessions"
)

var (
	manager = sessions.New(sessions.Config{Cookie: "iris_cookie"})
)

func main() {
	app := iris.New()
	app.Use(recover2.New())
	app.Logger().SetLevel("warn")
	app.Use(logger.New(logger.Config{}))

	app.StaticWeb("/static", "./static")

	app.Controller("/", new(controller.IndexController), manager)
	app.RegisterView(iris.HTML("./view", ".html"))

	// 1. GET  splash page
	//app.Handle("GET", "/", func(context context.Context) {
	//	context.HTML("<h1>Welcome Iris.</h1>")
	//})

	// 2. GET
	//app.Get("/string", func(context context.Context) {
	//	context.WriteString("Hello World.")
	//})

	// 3. GET
	//app.Get("/json", func(context context.Context) {
	//	context.JSON(iris.Map{"code": 200, "msg": "请求成功", "data": nil})
	//})

	// 4. template
	//app.Get("/index", func(context context.Context) {
	//	context.ViewData("msg", "HAHAHEHEHE")
	//	context.View("index.html")
	//})

	//注册路由
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
