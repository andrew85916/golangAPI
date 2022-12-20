package usecase

import (
	"crypto/sha256"
	"fmt"
	"golang_api/domain"
	errMessage "golang_api/user"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	User *domain.User `json:"user"`
	jwt.StandardClaims
}

type UserUsecase struct {
	repo           domain.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewUserUsecase(
	repo domain.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL time.Duration) *UserUsecase {

	return &UserUsecase{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}

// type UserUsecase interface {
// 	SignUP(username, password string) error
// 	SignIn(username, password string) (string, error)
// 	ParseToken(accessToken string) (User, error)
// }

func (u *UserUsecase) SignUp(username, password string) error {
	// Adding hash salt SHA256
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(u.hashSalt))

	user := &domain.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) SignIn(username, password string) (string, error) {

	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(u.hashSalt))
	user := &domain.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}
	user, err := u.repo.GetUser(user)

	if err != nil {
		return "", errMessage.ErrUserNotFound
	}
	claims := AuthClaims{
		User:           user,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: jwt.At(time.Now().Add(u.expireDuration)),
		},
	}

	//declared HS256 and add Payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// set signingKey as secret then generate jwt
	return token.SignedString(u.signingKey)
}

func (u *UserUsecase) ParseToken(accessToken string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}
	return nil, errMessage.ErrInvalidAccessToken
}
