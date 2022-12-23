package http

import (
	"fmt"
	"golang_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleGetInput struct {
	Aurhor  string `json:"author"`
	Content string `json:"content"`
}

type articleUpdateInput struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

type articleDeleteInput struct {
	Id int `json:"id"`
}

type articlePostInput struct {
	Content string `json:"content"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	usecase domain.ArticleUsecase
}

type getUsernameResponse struct {
	Username string `json:"username"`
}

func NewArticleHandler(r *gin.RouterGroup, us domain.ArticleUsecase) {
	handler := &ArticleHandler{
		usecase: us,
	}
	r.POST("/post_article", handler.PostArticle)
	r.GET("/get_author_articles", handler.GetArticleListByAuthor)
	r.GET("/get_self_articles", handler.GetSelfArticleList)
	r.GET("/get_others_articles", handler.GetOthersArticleList)
	r.GET("/get_username", handler.GetUsernameByToken)
	r.PUT("/update_article", handler.UpdateArticleContentById)
	r.DELETE("/delete_article", handler.DeleteArticleById)
}

func (ar *ArticleHandler) PostArticle(ctx *gin.Context) {
	arIn := new(articlePostInput)

	if err := ctx.BindJSON(arIn); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	author, exist := ctx.Get("username")
	authorStr := fmt.Sprintf("%v", author)
	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := ar.usecase.PostArticle(authorStr, arIn.Content); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)

}

func (ar *ArticleHandler) GetArticleListByAuthor(ctx *gin.Context) {
	arIn := new(articleGetInput)

	if err := ctx.BindJSON(arIn); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	articles, err := ar.usecase.GetArticleListByAuthor(arIn.Aurhor)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func (ar *ArticleHandler) GetSelfArticleList(ctx *gin.Context) {

	author, exist := ctx.Get("username")
	authorStr := fmt.Sprintf("%v", author)
	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	articles, err := ar.usecase.GetArticleListByAuthor(authorStr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func (ar *ArticleHandler) GetOthersArticleList(ctx *gin.Context) {

	author, exist := ctx.Get("username")
	authorStr := fmt.Sprintf("%v", author)
	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	articles, err := ar.usecase.GetOthersArticleList(authorStr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func (ar *ArticleHandler) GetUsernameByToken(ctx *gin.Context) {

	username, exist := ctx.Get("username")
	usernameStr := fmt.Sprintf("%v", username)
	if !exist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, getUsernameResponse{Username: usernameStr})
}

func (ar *ArticleHandler) UpdateArticleContentById(ctx *gin.Context) {
	arIn := new(articleUpdateInput)
	if err := ctx.BindJSON(arIn); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := ar.usecase.UpdateArticleContentById(arIn.Id, arIn.Content)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (ar *ArticleHandler) DeleteArticleById(ctx *gin.Context) {
	arIn := new(articleDeleteInput)
	// idStr := ctx.Param("id")
	// id, _ := strconv.Atoi(idStr)
	if err := ctx.BindJSON(arIn); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// err := ar.usecase.DeleteArticleById(arIn.Id)
	err := ar.usecase.DeleteArticleById(arIn.Id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}
