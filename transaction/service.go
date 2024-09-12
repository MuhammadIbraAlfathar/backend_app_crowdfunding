package transaction

import (
	"errors"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/campaign"
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/payment"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

type Service interface {
	GetTransactionsByCampaignId(input GetTransactionCampaignById) ([]Transaction, error)
	GetTransactionsByUserId(UserId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{
		repository, campaignRepository, paymentService,
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

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignId = input.CampaignId
	transaction.Status = "pending"
	transaction.UserId = input.User.ID

	newTransaction, err := s.repository.CreateTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		Id:     newTransaction.Id,
		Amount: newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl

	updatedTransaction, err := s.repository.UpdateTransaction(newTransaction)
	if err != nil {
		return updatedTransaction, err
	}

	return updatedTransaction, nil
}
