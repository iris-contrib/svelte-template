package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.HandleDir("/", iris.Dir("./app/public"))
	app.Listen(":8080")
}
