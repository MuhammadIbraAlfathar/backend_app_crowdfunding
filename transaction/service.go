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
	GetTransactionsByUserId(UserId int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{
		repository, campaignRepository,
	}
}

func (s *service) GetTransactionsByCampaignId(input GetTransactionCampaignById) ([]Transaction, error) {
	campaigns, err := s.campaignRepository.FindCampaignByCampaignId(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaigns.UserId != input.User.ID {
		return []Transaction{}, errors.New("error")
	}

	transactions, err := s.repository.GetTransactionByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserId(UserId int) ([]Transaction, error) {
	transactions, err := s.repository.GetTransactionByUserId(UserId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
