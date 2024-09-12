package transaction

import (
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/campaign"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"time"
)

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	PaymentUrl string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
