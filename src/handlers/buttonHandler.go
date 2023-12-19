package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/todo-htmx/src/views/components"
)

var (
	buttonCount = 0
)

func IncrementButtonCount(c echo.Context) error {
	buttonCount++
	count := strconv.Itoa(buttonCount)
	return Render(c, components.Button(count))
}