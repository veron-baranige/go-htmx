package main

import (
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/todo-htmx/src/routes"
)

func main() {
    e := echo.New()
    e.Static("/static", "static")
    routes.Setup(e)
    e.Logger.Fatal(e.Start(":8080"))
}
