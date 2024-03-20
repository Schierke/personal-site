package handler

import (
	"fmt"

	"github.com/Schierke/personal-site/model"
	"github.com/Schierke/personal-site/view/articles"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type ArticleRepository interface {
	GetArticleById(id string) (model.Article, error)
	GetArticles() ([]model.Article, error)
}

type ArticleHandler struct {
	repo ArticleRepository
}

func NewArticleHandler(repo ArticleRepository) *ArticleHandler {
	handler := &ArticleHandler{
		repo: repo,
	}

	return handler
}

func (h ArticleHandler) RegisterRoutes(router *echo.Echo) {
	router.GET("/blogs", h.ListArticles)
	router.GET("/articles/:id", h.ShowArticle)
}

func (h ArticleHandler) ListArticles(c echo.Context) error {
	list, err := h.repo.GetArticles()
	if err != nil {
		return err
	}
	return render(c, articles.ListArticles(list))
}

func (h ArticleHandler) ShowArticle(c echo.Context) error {
	id := c.Param("id")
	articleData, err := h.repo.GetArticleById(id)

	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("Error getting article: can't find article with %s", id))
	}

	return render(c, articles.ShowArticle(articleData))
}
