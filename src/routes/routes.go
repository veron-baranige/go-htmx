package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/todo-htmx/src/handlers"
)

func Setup(e *echo.Echo) {
	e.GET("/", handlers.RenderIndexPage)
	e.POST("/clicked", handlers.IncrementButtonCount)
}