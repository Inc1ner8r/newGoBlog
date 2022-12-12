package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/newGoBlog/app/db"
	"github.com/inciner8r/newGoBlog/app/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = db.ConnectDb()

func Register(c *gin.Context) {

	//// NON ORM CODE

	// var user models.User

	// if err := c.BindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	// 	return
	// }

	// if results, err := DB.Query(`INSERT INTO users(username, name, password) VALUES (? , ? , ?)`, user.Username, user.Name, user.Password); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	// 	return
	// } else {
	// 	fmt.Println(results.Columns())
	// }

	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 1)
	if err != nil {
		fmt.Println("hashing error")
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error})
		return
	}

	user.Password = string(hash)

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, user)

}

func GetUsers(c *gin.Context) {

	//// NON ORM CODE

	// if results, err := DB.Query(`SELECT * FROM users`); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	// 	return
	// } else {
	// 	fmt.Println(results.Columns())
	// 	users := []models.User{}
	// 	for results.Next() {
	// 		var user models.User

	// 		err = results.Scan(&user.ID, &user.Name, &user.Password, &user.Username)
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}

	// 		users = append(users, user)
	// 	}
	// 	fmt.Println(users[0])
	// }

	var users []models.User

	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {

}
