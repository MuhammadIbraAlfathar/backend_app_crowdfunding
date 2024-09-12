package payment

import (
	"github.com/MuhammadIbraAlfathar/backend_app_crowdfunding/user"
	"github.com/veritrans/go-midtrans"
	"os"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	midtransClient := midtrans.NewClient()
	clientKey := os.Getenv("CLIENT_KEY_MIDTRANS")
	serverKey := os.Getenv("SERVER_KEY_MIDTRANS")
	midtransClient.ServerKey = serverKey
	midtransClient.ClientKey = clientKey
	midtransClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midtransClient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
