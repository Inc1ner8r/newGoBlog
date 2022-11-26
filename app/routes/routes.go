package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "api is up"})
	})
}
