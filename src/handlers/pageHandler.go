package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/todo-htmx/src/views"
	"github.com/veron-baranige/todo-htmx/src/views/components"
)

func RenderIndexPage(c echo.Context) error {
	return Render(c, views.IndexPage(components.Button("0")))
}