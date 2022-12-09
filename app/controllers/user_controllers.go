package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/db"
	"github.com/inciner8r/newGoBlog/app/models"
)

var DB = db.DB

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	user_int, _ := json.Marshal(user)
	user_string := string(user_int)
	fmt.Println(user_string)
	if sql_rows, err := DB.Query(`INSERT INTO users(name, username, password) VALUES (? , ? , ?)`, user.Name, user.Username, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	} else {
		fmt.Println(sql_rows)
	}

}
