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

func (a *ArticleUsecase) PostArticle(author, content string) error {

	article := &domain.Article{
		Author:    author,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := a.articleRepo.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil

}

// func (u *UserUsecase) SignUp(username, password string) error {
// 	// Adding hash salt SHA256
// 	pwd := sha256.New()
// 	pwd.Write([]byte(password))
// 	pwd.Write([]byte(u.hashSalt))

// 	user := &domain.User{
// 		Username: username,
// 		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
// 	}
// 	return u.repo.CreateUser(user)
// }
