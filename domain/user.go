package domain

// type User struct {
// 	ID       string `json:"id"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }
// `gorm:"primary_key:auto_increment" json:"id"`
type User struct {
	ID       string `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Password string `gorm:"->;<-;not null" json:"password"`
}

type UserRepository interface {
	GetUserList(map[string]interface{}) ([]*User, error)
	GetUser(user *User) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}

//
type UserUsecase interface {
	SignUp(username, password string) error
	SignIn(username, password string) (string, error)
	ParseToken(accessToken string) (*User, error)
}
