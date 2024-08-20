package handler

import (
	"fmt"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/auth"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/helper"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		userService: userService,
		authService: authService,
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

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatterUser := user.FormatterUser(newUser, token)

	response := helper.ResponseApi(messageSuccess, statusSuccess, http.StatusOK, formatterUser)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {

	messageSuccess := "Login Success"
	statusSuccess := "success"
	messageFailed := "Login Failed"
	statusFailed := "error"

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	formatterUser := user.FormatterUser(loginUser, token)

	response := helper.ResponseApi(messageSuccess, statusSuccess, http.StatusOK, formatterUser)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) IsEmailAvailable(c *gin.Context) {
	var input user.CheckEmailInput

	statusSuccess := "success"
	messageFailed := "Email Checking Failed"
	statusFailed := "error"

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	available, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ResponseApi(messageFailed, statusFailed, http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": available,
	}

	metaMessage := "Email Has Been Registered"
	if available {
		metaMessage = "Email is available"
	}

	response := helper.ResponseApi(metaMessage, statusSuccess, http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	messageFailed := "Failed to upload avatar image"
	messageSuccess := "Avatar successfuly uploaded"
	statusError := "error"
	statusSuccess := "success"

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.ResponseApi(messageFailed, statusError, http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Harusnya dapat dari jwt
	userId := 16

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.ResponseApi(messageFailed, statusError, http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.ResponseApi(messageFailed, statusError, http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}

	response := helper.ResponseApi(messageSuccess, statusSuccess, http.StatusOK, data)
	c.JSON(http.StatusOK, response)
	return
}
