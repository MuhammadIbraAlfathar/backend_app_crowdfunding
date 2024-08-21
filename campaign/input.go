package campaign

type GetDetailCampaignInput struct {
	Id int `uri:"id" binding:"required"`
}
