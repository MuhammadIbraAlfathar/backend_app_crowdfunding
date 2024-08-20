package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAllCampaign() ([]Campaign, error)
	FindCampaignByUserId(userId int) ([]Campaign, error)
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
	err := r.db.Find(campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindCampaignByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userId).Find(campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
