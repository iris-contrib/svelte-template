package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.HandleDir("/", "./app/public")

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
