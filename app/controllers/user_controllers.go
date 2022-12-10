package controllers

import (
	"github.com/gin-gonic/gin"
)

// var DB = db.DB

func CreateUser(c *gin.Context) {

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
}
