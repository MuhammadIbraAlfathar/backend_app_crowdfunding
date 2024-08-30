package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAllCampaign() ([]Campaign, error)
	FindCampaignByUserId(userId int) ([]Campaign, error)
	FindCampaignByCampaignId(campaignId int) (Campaign, error)
	SaveCampaign(campaign Campaign) (Campaign, error)
	UpdateCampaign(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAllCampaign() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindCampaignByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindCampaignByCampaignId(campaignId int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", campaignId).Preload("User").Preload("CampaignImages").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) SaveCampaign(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) UpdateCampaign(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, err
}
