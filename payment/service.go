package payment

import (
	user "rexencorpstartup/User"
	"strconv"

	// "github.com/midtrans/midtrans-go"
	// "github.com/midtrans/midtrans-go/snap"
	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service{
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) { 
	// midtrans.ServerKey = ""
	// midtrans.Environment = midtrans.Sandbox
	
	// s := snap.Client
	// s.New("gblk", midtrans.Sandbox)
	
	
	midclient := midtrans.NewClient()
	midclient.ServerKey = "hehekampang"
	midclient.ClientKey = "hehekampang"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client : midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail : &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil{
		return "", err 
	}
	return snapTokenResp.RedirectURL, nil

}