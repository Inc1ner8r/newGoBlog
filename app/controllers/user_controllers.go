package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/inciner8r/newGoBlog/app/db"
	"github.com/inciner8r/newGoBlog/app/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = db.ConnectDb()

var key = []byte("key")

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

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("hashing error")
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
	// 			panic(err.Error())w
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
	var credentials models.Credentials
	var expected models.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}

	if err := DB.Table("users").Where("username = ?", credentials.Username).First(&expected).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(expected.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 12)

	claims := &models.Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedtoken, err := token.SignedString(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		fmt.Println("ok")
		return
	}
	c.SetCookie("jwt", signedtoken, int(time.Hour)*12/int(time.Second), "/", "localhost", false, true)
}

func ValidateJWT(c *gin.Context) string {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "cookie not found"})
		log.Fatal("unauthorized")
	}
	token, err := jwt.ParseWithClaims(cookie, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		log.Fatal("unauthorized")
	}
	claims := token.Claims.(*models.Claims)
	return claims.Username
}
