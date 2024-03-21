package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Schierke/personal-site/model"
	"gopkg.in/yaml.v3"
)

type articleRepository struct {
}

func NewArticleRepository() *articleRepository {
	return &articleRepository{}
}

// Getting articles from articles directory
// the article are stored each inside a directory, which is the article adress
// There are 2 files inside each directory, the first one is the article content (post.md)
// the other is the article metadata (meta.yml)
func (r articleRepository) GetArticles() ([]model.Article, error) {
	dir, err := os.ReadDir("./articles")
	if err != nil {
		return nil, err
	}

	var articles []model.Article

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}

		articleDir := filepath.Join("articles", entry.Name())

		metadata, err := getArticleMetadata(articleDir)

		if err != nil {
			return nil, err
		}

		article := model.Article{
			ID:   entry.Name(),
			Meta: *metadata,
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// getting article with id from articles directory
func (r articleRepository) GetArticleById(id string) (model.Article, error) {
	articleDir := filepath.Join("articles", id)

	metadata, err := getArticleMetadata(articleDir)
	if err != nil {
		return model.Article{}, err
	}

	content, err := getArticleContent(articleDir)
	if err != nil {
		return model.Article{}, err
	}

	article := model.Article{
		ID:      id,
		Meta:    *metadata,
		Content: content,
	}

	return article, nil
}

// getting article metadata from meta.yml file
func getArticleMetadata(path string) (*model.ArticleMetadata, error) {
	metaFile := filepath.Join(path, "meta.yml")
	file, err := os.ReadFile(metaFile)

	if err != nil {
		return nil, err
	}

	var metadata model.ArticleMetadata
	err = yaml.Unmarshal(file, &metadata)
	if err != nil {
		return nil, err
	}

	metadata.Link = fmt.Sprintf("/blog/%s", filepath.Base(path))

	return &metadata, nil
}

// getting article content from post.md file
func getArticleContent(path string) ([]byte, error) {
	postFile := filepath.Join(path, "post.md")
	file, err := os.ReadFile(postFile)

	if err != nil {
		return nil, err
	}

	return file, nil
}
