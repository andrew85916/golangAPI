package http

import (
	"golang_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleInput struct {
	Aurhor  string `json:"author"`
	Content string `json:"content"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	usecase domain.ArticleUsecase
}

func NewArticleHandler(r *gin.RouterGroup, us domain.ArticleUsecase) {
	handler := &ArticleHandler{
		usecase: us,
	}
	// r.GET("/articles", handler.GetArticleListByUser)
	r.POST("/post_article", handler.PostArticle)
	r.GET("/get_article", handler.GetArticleListByAuthor)
}

func (ar *ArticleHandler) PostArticle(ctx *gin.Context) {
	arIn := new(articleInput)
	if err := ctx.BindJSON(arIn); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := ar.usecase.PostArticle(arIn.Aurhor, arIn.Content); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)

}

func (ar *ArticleHandler) GetArticleListByAuthor(ctx *gin.Context) {
	arIn := new(articleInput)
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
