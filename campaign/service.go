package campaign

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignByCampaignId(input GetDetailCampaignInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputId GetDetailCampaignInput, inputData CreateCampaignInput) (Campaign, error)
	CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) UpdateCampaign(inputId GetDetailCampaignInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindCampaignByCampaignId(inputId.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.ID {
		return campaign, errors.New("failed updated campaign")
	}

	campaign.Name = inputData.Name
	campaign.Description = inputData.Description
	campaign.ShortDescription = inputData.ShortDescription
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	updatedCampaign, err := s.repository.UpdateCampaign(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) CreateCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindCampaignByCampaignId(input.CampaignId)
	if err != nil {
		return CampaignImage{}, err
	}
	if campaign.UserId != input.User.ID {
		return CampaignImage{}, errors.New("failed upload campaign image")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignId)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignId = input.CampaignId
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	image, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return campaignImage, err
	}

	return image, nil
}
