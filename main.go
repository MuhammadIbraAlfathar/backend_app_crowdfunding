package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database")

	//var users []user.User
	//
	//db.Find(&users)
	//
	//for _, u := range users {
	//	fmt.Println(u.Name)
	//	fmt.Println(u.Email)
	//}
}
