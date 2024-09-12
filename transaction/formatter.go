package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransactionCampaign(transaction Transaction) CampaignTransactionFormatter {
	transactionFormatter := CampaignTransactionFormatter{}
	transactionFormatter.Id = transaction.Id
	transactionFormatter.Name = transaction.User.Name
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.CreatedAt = transaction.CreatedAt

	return transactionFormatter
}

func FormatTransactionsCampaigns(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		transactionFormatter := FormatTransactionCampaign(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)

	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	Id        int               `json:"id"`
	Status    string            `json:"status"`
	Amount    int               `json:"amount"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Status = transaction.Status
	formatter.Amount = transaction.Amount

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) < 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		userTransaction := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransaction)
	}

	return userTransactionsFormatter
}

type CreateTransactionResponseFormatter struct {
	Id         int    `json:"id"`
	CampaignId int    `json:"campaign_id"`
	UserId     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentUrl string `json:"payment_url"`
}

func FormatterCreateTransactionResponse(transaction Transaction) CreateTransactionResponseFormatter {
	response := CreateTransactionResponseFormatter{}
	response.Id = transaction.Id
	response.CampaignId = transaction.CampaignId
	response.UserId = transaction.UserId
	response.Amount = transaction.Amount
	response.Status = transaction.Status
	response.Code = transaction.Code
	response.PaymentUrl = transaction.PaymentUrl

	return response
}
