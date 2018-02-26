package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type MainController struct {
}

func (c MainController) Init(g *echo.Group) {
	g.GET("", c.goIndex)
}

func (MainController) goIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
