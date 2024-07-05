package handler

import (
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/helper"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	messageSuccess := "Account has been registered"
	statusSuccess := "success"
	messageFailed := "Register Account Failed"
	statusFailed := "error"

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		log.Println(err)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		log.Println(err)
		return
	}

	token := "testing"
	formatterUser := user.FormatterUser(newUser, token)

	response := helper.ResponseApi(messageSuccess, statusSuccess, http.StatusOK, formatterUser)

	c.JSON(http.StatusOK, response)
}
