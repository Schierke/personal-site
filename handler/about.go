package handler

import (
	"github.com/Schierke/personal-site/view/about"
	"github.com/labstack/echo/v4"
)

type AboutHandler struct {
}

func NewAboutHandler() *AboutHandler {
	handler := &AboutHandler{}

	return handler
}

func (h AboutHandler) RegisterRoutes(router *echo.Echo) {
	router.GET("/about", h.AboutShow)
}

func (h AboutHandler) AboutShow(c echo.Context) error {
	return render(c, about.Show())
}
