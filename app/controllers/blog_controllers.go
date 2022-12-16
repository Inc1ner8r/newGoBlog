package controllers

import (
	"github.com/gin-gonic/gin"
)

func PostBog(c *gin.Context) {
	ValidateJWT(c)
}
