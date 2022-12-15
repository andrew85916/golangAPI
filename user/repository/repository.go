package repository

import (
	"golang_api/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	orm *gorm.DB
}

func NewUserRepository(orm *gorm.DB) domain.UserRepository {
	return &UserRepository{
		orm: orm,
	}
}

// type UserRepository interface {
// 	GetUserList(map[string]interface{}) ([]*User, error)
// 	GetUser(u *User) (*User, error)
// 	CreateUser(u *User) (*User, error)
// 	UpdateUser(u *User, data map[string]interface{}) (*User, error)
// 	DeleteUser(u *User) error
// }

func (u *UserRepository) GetUserList(users map[string]interface{}) ([]*domain.User, error) {

	userList := make([]*domain.User, 0)
	err := u.orm.Find(userList, users).Error
	return userList, err
}
func (u *UserRepository) GetUser(user *domain.User) (*domain.User, error) {

	err := u.orm.Where(map[string]interface{}{"username": user.Username, "password": user.Password}).Take(&user).Error
	return user, err

}
func (u *UserRepository) CreateUser(user *domain.User) error {
	err := u.orm.Create(&user).Error

	return err
}
func (u *UserRepository) UpdateUser(user *domain.User) error {
	err := u.orm.Save(&user).Error

	return err
}
func (u *UserRepository) DeleteUser(user *domain.User) error {
	err := u.orm.Delete(&user).Error
	return err
}
