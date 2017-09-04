package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"studio.xiwi/controllers"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/flash", controllers.Flash)

	app.Run(iris.Addr(":8087"), iris.WithoutServerError(iris.ErrServerClosed))
}
