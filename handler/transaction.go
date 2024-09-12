package handler

import (
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/helper"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/transaction"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{
		service,
	}
}

func (h *transactionHandler) GetTransactionsCampaignByCampaignId(c *gin.Context) {
	var input transaction.GetTransactionCampaignById
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ResponseApi("Failed to get transactions campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignId(input)
	if err != nil {
		response := helper.ResponseApi("Failed to get transactions campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseApi("Success to get transactions campaign", "success", http.StatusOK, transaction.FormatTransactionsCampaigns(transactions))
	c.JSON(http.StatusOK, response)
	return

}

func (h *transactionHandler) GetTransactionsByUserId(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserId(userId)
	if err != nil {
		response := helper.ResponseApi("Failed to get transactions campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseApi("Success to get transactions campaign", "success", http.StatusOK, transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
	return
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ResponseApi("Failed to create transaction", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.ResponseApi("Failed to create transaction", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		log.Println(err.Error())
		return
	}

	response := helper.ResponseApi("Success to create transaction", "success", http.StatusOK, transaction.FormatterCreateTransactionResponse(newTransaction))
	c.JSON(http.StatusOK, response)
	return

}
