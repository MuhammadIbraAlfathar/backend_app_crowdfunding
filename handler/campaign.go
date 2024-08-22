package handler

import (
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/campaign"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/helper"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.ResponseApi("Error to get campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseApi("Success to get campaign", "success", http.StatusOK, campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetDetailCampaignById(c *gin.Context) {
	var input campaign.GetDetailCampaignInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ResponseApi("Failed to get detail campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByCampaignId(input)
	if err != nil {
		response := helper.ResponseApi("Failed to get detail campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseApi("Success get detail campaign", "success", http.StatusOK, campaign.FormatDetailCampaign(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ResponseApi("Failed to create campaign", "error", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		log.Println(err)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.ResponseApi("Failed to create campaign", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseApi("Success create campaign", "success", http.StatusOK, campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}
