package model

type ArticleMetadata struct {
	Title            string
	Link             string
	Date             string
	ShortDescription string
	Tags             []string
}

type Article struct {
	ID      string
	Content []byte
	Meta    ArticleMetadata
}
