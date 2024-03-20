package handler

import (
	"github.com/Schierke/personal-site/view/index"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h IndexHandler) RegisterRoutes(router *echo.Echo) {
	router.GET("/", h.IndexShow)
}

func (h IndexHandler) IndexShow(c echo.Context) error {
	return render(c, index.ShowIndex())
}
