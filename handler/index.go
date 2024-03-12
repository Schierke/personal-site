package handler

import (
	"github.com/Schierke/personal-site/model"
	"github.com/Schierke/personal-site/view/index"
	"github.com/labstack/echo/v4"
)

type IndexRepositoy interface {
	GetArticles() ([]model.Article, error)
}

type IndexHandler struct {
	repo IndexRepositoy
}

func NewIndexHandler(repo IndexRepositoy) *IndexHandler {
	handler := &IndexHandler{
		repo: repo,
	}

	return handler
}

func (h IndexHandler) RegisterRoutes(router *echo.Echo) {
	router.GET("/", h.IndexShow)
}

func (h IndexHandler) IndexShow(c echo.Context) error {
	articles, err := h.repo.GetArticles()
	if err != nil {
		return err
	}
	return render(c, index.Show(articles))
}
