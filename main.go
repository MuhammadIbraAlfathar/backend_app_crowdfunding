package main

import (
	"fmt"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/auth"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/campaign"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/handler"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/helper"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/payment"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/transaction"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/startup_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database")

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	//auth
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/check-email", userHandler.IsEmailAvailable)

	api.Use(authMiddleware(authService, userService))
	{
		//campaign
		api.GET("/campaigns", campaignHandler.GetCampaigns)
		api.GET("/campaign/:id", campaignHandler.GetDetailCampaignById)
		api.POST("/campaign", campaignHandler.CreateCampaign)
		api.PUT("/campaign/:id", campaignHandler.UpdateCampaign)
		api.POST("campaign/image", campaignHandler.UploadImage)

		//upload avatar
		api.POST("/avatar", userHandler.UploadAvatar)

		//transactions
		api.GET("/transactions/campaign/:id", transactionHandler.GetTransactionsCampaignByCampaignId)
		api.GET("/transactions/campaign/user", transactionHandler.GetTransactionsByUserId)
		api.POST("/transactions", transactionHandler.CreateTransaction)
	}

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponseApi("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer data token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ResponseApi("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, oke := token.Claims.(jwt.MapClaims)
		if !oke || !token.Valid {
			response := helper.ResponseApi("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		userData, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.ResponseApi("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userData)

	}
}
