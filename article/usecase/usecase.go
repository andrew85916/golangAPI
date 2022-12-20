package usecase

import (
	"golang_api/domain"
	"time"
)

type ArticleUsecase struct {
	articleRepo domain.ArticleRepository
}

// Create a new articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleUsecase(repo domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{
		articleRepo: repo,
	}
}

func (ar *ArticleUsecase) GetArticleListByAuthor(author string) ([]*domain.Article, error) {
	articles, err := ar.articleRepo.GetArticleListByAuthor(author)

	return articles, err

}

func (ar *ArticleUsecase) GetOthersArticleList(author string) ([]*domain.Article, error) {
	articles, err := ar.articleRepo.GetOthersArticleList(author)

	return articles, err
}

func (ar *ArticleUsecase) PostArticle(author, content string) error {

	article := &domain.Article{
		Author:    author,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := ar.articleRepo.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil

}

func (ar *ArticleUsecase) DeleteArticleById(id int) error {
	err := ar.articleRepo.DeleteArticleById(id)
	return err
}
