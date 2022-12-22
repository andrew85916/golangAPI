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

	err := a.orm.Order("id desc").Where("author=?", author).Find(&articles).Error

	return articles, err
}

func (a *articleRepository) GetOthersArticleList(author string) ([]*domain.Article, error) {
	articles := make([]*domain.Article, 0)
	err := a.orm.Order("id desc").Not("author=?", author).Find(&articles).Error
	return articles, err
}

func (a *articleRepository) UpdateArticleContent(article *domain.Article) error {

	// err := a.orm.Model(&article).Updates(map[string]interface{}{"content": content, "updated_at": updateAt}).Error
	err := a.orm.Debug().Model(&article).Updates(map[string]interface{}{"content": article.Content, "updated_at": article.UpdatedAt}).Error
	// err := a.orm.Model(&article).Updates(map[string]interface{}{"content": article.Content, "updated_at": article.UpdatedAt}).Error
	return err
}
func (a *articleRepository) DeleteArticleById(id int) error {

	article := new(domain.Article)

	err := a.orm.Where("id=?", id).Delete(&article).Error
	return err
}
