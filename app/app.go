package app

import (
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/routes"
)

func Init() {
	r := gin.Default()
	routes.SetRoutes(r)
	r.Run(":3000")
}
