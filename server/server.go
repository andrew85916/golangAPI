package server

import (
	"fmt"
	"golang_api/domain"

	_articleHttp "golang_api/article/delivery/http"
	_articleRepo "golang_api/article/repository"
	_articleUsecase "golang_api/article/usecase"
	_userHttp "golang_api/user/delivery/http"
	_userMiddleware "golang_api/user/delivery/http/middleware"
	_userRepo "golang_api/user/repository"
	_userUsecase "golang_api/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	userUC    domain.UserUsecase
	articleUC domain.ArticleUsecase
}

func NewApp() *App {
	db := initDB()
	userRepo := _userRepo.NewUserRepository(db)
	articleRepo := _articleRepo.NewArticleRepository(db)
	return &App{
		articleUC: _articleUsecase.NewArticleUsecase(articleRepo),
		userUC: _userUsecase.NewUserUsecase(
			userRepo,
			viper.GetString(`auth.hash_salt`),
			[]byte(viper.GetString(`auth.signing_key`)),
			viper.GetDuration(`auth.token_ttl`),
		),
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func (a *App) Run() {
	// Init gin handler
	router := gin.Default()
	// Cross-Origin Resource Sharing

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		CORSMiddleware(),
	)
	// Set up http handlers
	// SignUp/SignIn endpoints
	_userHttp.NewUserHandler(router, a.userUC)

	// API endpoints
	authMiddleware := _userMiddleware.NewAuthMiddleware(a.userUC)
	api := router.Group("/api", authMiddleware)
	// api := router.Group("/api")

	_articleHttp.NewArticleHandler(api, a.articleUC)

	router.Run()
}

func initDB() *gorm.DB {
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.password`)
	dbHost := viper.GetString(`database.host`)
	dbName := viper.GetString(`database.name`)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to databasee")

	}
	return db
}
