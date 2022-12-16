package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/models"
)

func PostBlog(c *gin.Context) {
	var user models.User

	var blog models.Blog
	username := ValidateJWT(c)
	if err := DB.Table("users").Where(`username = ?`, username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}
	blog.Author_id = user.Id

	if err := DB.Create(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func DisplayAllBlogs(c *gin.Context) {
	ValidateJWT(c)
	var blogs []models.Blog
	if err := DB.Find(&blogs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, blogs)

}
