package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sranjan001/bookstore_oauth-api/src/domain/access_token"
	"github.com/sranjan001/bookstore_oauth-api/src/http"
	"github.com/sranjan001/bookstore_oauth-api/src/respository/db"
	"github.com/sranjan001/bookstore_oauth-api/src/respository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository(), rest.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8081")
}
