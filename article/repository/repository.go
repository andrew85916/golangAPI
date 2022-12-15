package repository

import (
	"golang_api/domain"

	"gorm.io/gorm"
)

type articleRepository struct {
	orm *gorm.DB
}

func NewArticleRepository(orm *gorm.DB) domain.ArticleRepository {
	return &articleRepository{
		orm: orm,
	}
}

// type ArticleRepository interface {
// 	GetArticleList(map[string]interface{}) ([]*Article, error)
// 	GetArticle(article *Article) (*Article, error)
// 	CreateArticle(article *Article) error
// 	UpdateArticle(article *Article) error
// 	DeleteArticle(article *Article) error
// }

func (a *articleRepository) CreateArticle(article *domain.Article) error {

	err := a.orm.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) GetArticleListByAuthor(author string) ([]*domain.Article, error) {
	articles := make([]*domain.Article, 0)
	// articleList := make([]*domain.Article, 0)
	// err := a.orm.Where("author=?", author).Find(articlesList,artic).Error
	err := a.orm.Where("author=?", author).Find(&articles).Error
	// articles := make([]*domain.Article, 0)
	// result := a.orm.Find(&articles)
	return articles, err
}

// func (a *articleRepository) GetArticleListByUser(articles map[string]interface{}) ([]*domain.Article, error) {

// 	articleList := make([]*domain.Article, 0)
// 	err := a.orm.Find(articleList, articles).Error
// 	return articleList, err
// }

// func (a *ArticleRepository) GetArticle(article *domain.Article) (*domain.Article, error) {
// 	err := a.orm.Save(&article).Error
// 	return err
// }

func (a *articleRepository) UpdateArticle(article *domain.Article) error {
	err := a.orm.Save(&article).Error
	return err
}
func (a *articleRepository) DeleteArticle(article *domain.Article) error {
	err := a.orm.Save(&article).Error
	return err
}
