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

	//userRepository := user.NewRepository(db)
	//userService := user.NewService(userRepository)
	//
	//userInput := user.RegisterUserInput{}
	//userInput.Email = "test@gmail.com"
	//userInput.Name = "testttt"
	//userInput.Password = "asras"
	//userInput.Occupation = "programmer"
	//
	//userService.RegisterUser(userInput)

	//var users []user.User
	//
	//db.Find(&users)
	//
	//for _, u := range users {
	//	fmt.Println(u.Name)
	//	fmt.Println(u.Email)
	//}
}
