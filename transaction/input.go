package transaction

import "github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"

type GetTransactionCampaignById struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}
