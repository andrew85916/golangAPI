package middleware

import (
	"golang_api/domain"
	errMessage "golang_api/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	usecase domain.UserUsecase
}

func NewAuthMiddleware(us domain.UserUsecase) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: us,
	}).AuthorizeJWT
}

func (m *AuthMiddleware) AuthorizeJWT(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == errMessage.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		ctx.AbortWithStatus(status)
		return
	}

	// Store user.Username info into Context
	// Get Username info from ctx.Get("username")
	ctx.Set("username", user.Username)
	ctx.Next()
}
