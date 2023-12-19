package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	c.Response().Header().Add("Content-Type", "text/html")
	return component.Render(c.Request().Context(), c.Response())
}