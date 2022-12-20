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

func (a *articleRepository) CreateArticle(article *domain.Article) error {

	err := a.orm.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) GetArticleListByAuthor(author string) ([]*domain.Article, error) {
	articles := make([]*domain.Article, 0)

	err := a.orm.Where("author=?", author).Find(&articles).Error

	return articles, err
}

func (a *articleRepository) GetOthersArticleList(author string) ([]*domain.Article, error) {
	articles := make([]*domain.Article, 0)
	err := a.orm.Not("author=?", author).Find(&articles).Error
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
func (a *articleRepository) DeleteArticleById(id int) error {
	article := new(domain.Article)

	err := a.orm.Where("id=?", id).Delete(&article).Error
	return err
}
