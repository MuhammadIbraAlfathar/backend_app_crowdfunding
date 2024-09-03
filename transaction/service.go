package transaction

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/campaign"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignId(input GetTransactionCampaignById) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{
		repository, campaignRepository,
	}
}

func (s *service) GetTransactionsByCampaignId(input GetTransactionCampaignById) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindCampaignByCampaignId(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("error")
	}

	transactions, err := s.repository.GetTransactionByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
