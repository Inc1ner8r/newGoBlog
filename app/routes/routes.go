package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/controllers"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "api is up"})
	})
	router.POST("/createUser", controllers.Register)
	router.GET("/getUsers", controllers.GetUsers)
	router.POST("/login", controllers.LoginTemp)
	router.POST("/postBlog", controllers.PostBlog)
	router.GET("/getBlogs", controllers.DisplayAllBlogs)
}
