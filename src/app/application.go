package app

import (
	"github.com/Kungfucoding23/bookstore_oauth_api/src/domain/access_token"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/http"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.Run()
}