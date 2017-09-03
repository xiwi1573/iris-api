package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<b>welcome to to iris web framework...</b>")
	})

	app.Run(iris.Addr(":8087"), iris.WithoutServerError(iris.ErrServerClosed))
}
