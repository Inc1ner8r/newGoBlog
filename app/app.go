package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/routes"
)

func Init() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.SetRoutes(r)
	r.Run(":3000")
}
