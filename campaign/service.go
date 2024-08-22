package campaign

import (
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignByCampaignId(input GetDetailCampaignInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindCampaignByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAllCampaign()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByCampaignId(input GetDetailCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindCampaignByCampaignId(input.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.SaveCampaign(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
