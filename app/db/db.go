package db

import (
	"fmt"
	"log"
	"os"

	"github.com/inciner8r/newGoBlog/app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func ConnectDb() *sql.DB {

// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	username := os.Getenv("MYSQL_USER")
// 	password := os.Getenv("MYSQL_PASSWORD")
// 	dbName := os.Getenv("MYSQL_DATABASE")
// 	DbString := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbName + "?parseTime=true"

// 	if db, err := sql.Open("mysql", DbString); err != nil {
// 		fmt.Println("nok")
// 		panic(err.Error())
// 	} else {
// 		fmt.Println("ok")
// 		//defer db.Close()
// 		return db
// 	}
// }

// var DB = ConnectDb()

func ConnectDb() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbName + "?parseTime=true"
	fmt.Println(dsn)
	// dsn := "root:a@tcp(127.0.0.1:3306)/goBlog?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{})
	//db.AutoMigrate(&models.Blog{})
	fmt.Println("db init")

	return db
}
