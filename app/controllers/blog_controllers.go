package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/db"
)

var dB = db.ConnectDb()

func GetBlogs(c *gin.Context) {
	dB
}
