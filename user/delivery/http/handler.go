package http

import (
	"golang_api/domain"
	errMessage "golang_api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(router *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		usecase: us,
	}
	router.POST("/user/sign_in", handler.SignIn)
	router.POST("/user/sign_up", handler.SignUp)

}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type signInResponse struct {
	Token string `json:"token"`
}

func (u *UserHandler) SignIn(ctx *gin.Context) {
	in := new(signInput)

	if err := ctx.BindJSON(in); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := u.usecase.SignIn(in.Username, in.Password)
	if err != nil {
		if err == errMessage.ErrUserNotFound {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, signInResponse{Token: token})
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	in := new(signInput)

	if err := ctx.BindJSON(in); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := u.usecase.SignUp(in.Username, in.Password); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
