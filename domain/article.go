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
	// GetArticleListByUser(map[string]interface{}) ([]*Article, error)
	// GetArticle(article *Article) (*Article, error)
	CreateArticle(article *Article) error
	UpdateArticle(article *Article) error
	DeleteArticle(article *Article) error
}

type ArticleUsecase interface {
	GetArticleListByAuthor(author string) ([]*Article, error)
	PostArticle(author string, content string) error
}

// type User struct {
// 	ID       string `gorm:"primary_key:auto_increment" json:"id"`
// 	Username string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
// 	Password string `gorm:"->;<-;not null" json:"password"`
// }

// type UserRepository interface {
// 	GetUserList(map[string]interface{}) ([]*User, error)
// 	GetUser(user *User) (*User, error)
// 	CreateUser(user *User) error
// 	UpdateUser(user *User) error
// 	DeleteUser(user *User) error
// }

// //
// type UserUsecase interface {
// 	SignUp(username, password string) error
// 	SignIn(username, password string) (string, error)
// 	ParseToken(accessToken string) (*User, error)
// }
