package domain

import "time"

type Article struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content" validation:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type ArticleRepository interface {
	GetArticleListByAuthor(author string) ([]*Article, error)
	GetOthersArticleList(author string) ([]*Article, error)
	CreateArticle(article *Article) error
	UpdateArticleContent(article *Article) error
	DeleteArticleById(id int) error
}

type ArticleUsecase interface {
	GetArticleListByAuthor(author string) ([]*Article, error)
	GetOthersArticleList(author string) ([]*Article, error)
	PostArticle(author string, content string) error
	UpdateArticleContentById(id int64, content string) error
	DeleteArticleById(id int) error
}
