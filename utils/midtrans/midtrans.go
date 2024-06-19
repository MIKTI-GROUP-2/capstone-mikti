package midtrans

import (
	"capstone-mikti/configs"
	"errors"
	"fmt"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"github.com/sirupsen/logrus"
)

type ChargeResponse = coreapi.ChargeResponse

type MidtransService interface {
	GenerateTransaction(result int, paymentType string) (*ChargeResponse, map[string]interface{}, error)
	TransactionStatus(notificationPayload map[string]interface{}) (int, string, error)
}

type midtransService struct {
	core coreapi.Client
}

func InitMidtrans(c *configs.ProgrammingConfig) MidtransService {
	var core coreapi.Client
	var env midtrans.EnvironmentType
	midtransEnv := c.MidtransEnv

	if strings.ToLower(midtransEnv) == "production" {
		env = midtrans.Production
	} else {
		env = midtrans.Sandbox
	}

	core.New(c.MidtransServerKey, env)

	return &midtransService{
		core: core,
	}
}

func (ms *midtransService) GenerateTransaction(result int, paymentType string) (*ChargeResponse, map[string]interface{}, error) {
	var chargeReq *coreapi.ChargeReq
	response := map[string]any{}

	if paymentType == "qris" {
		chargeReq = &coreapi.ChargeReq{
			PaymentType: "qris",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  "Q-" + example.Random(),
				GrossAmt: int64(result),
			},
		}
	}

	if paymentType == "gopay" {
		chargeReq = &coreapi.ChargeReq{
			PaymentType: "gopay",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  "G-" + example.Random(),
				GrossAmt: int64(result),
			},
		}
	}

	if paymentType == "bca" || paymentType == "bni" || paymentType == "bri" {
		var midtransBank midtrans.Bank

		switch paymentType {
		case "bca":
			midtransBank = midtrans.BankBca
		case "bni":
			midtransBank = midtrans.BankBni
		case "bri":
			midtransBank = midtrans.BankBri
		default:
			midtransBank = midtrans.BankBca
		}

		chargeReq = &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtransBank,
			},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  "B-" + example.Random(),
				GrossAmt: int64(result),
			},
		}
	}

	chargeResp, err := ms.core.ChargeTransaction(chargeReq)

	if err != nil {
		logrus.Error("ERROR : Charge Transaction Midtrans : ", err.Error())
		return nil, nil, err
	}

	if paymentType == "qris" || paymentType == "gopay" {
		if len(chargeResp.Actions) > 0 {
			for _, action := range chargeResp.Actions {
				switch action.Name {
				case "generate-qr-code":
					deepLinkUrl := action.URL
					response["callback_url"] = deepLinkUrl
					response["payment_type"] = "qris"
				case "deeplink-redirect":
					deepLinkUrl := action.URL
					response["callback_url"] = deepLinkUrl
					response["payment_type"] = "gopay"
				}
			}
		}
	}

	if paymentType == "bca" || paymentType == "bni" || paymentType == "bri" {
		var vaAccount string
		for _, va := range chargeResp.VaNumbers {
			if va.Bank == paymentType {
				vaAccount = va.VANumber
				break
			}
		}
		response["payment_type"] = paymentType
		response["va_account"] = vaAccount
	}

	return chargeResp, response, nil
}

func (ms *midtransService) TransactionStatus(notificationPayload map[string]interface{}) (int, string, error) {
	var paymentStatus int
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return 0, "", errors.New("ERROR Order ID Not Found")
	}

	transactionStatusResp, e := ms.core.CheckTransaction(orderId)
	if e != nil {
		return 0, "", errors.New("ERROR Order ID Not Found")
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					fmt.Println("Payment status challenged")
					paymentStatus = 1
					return paymentStatus, transactionStatusResp.OrderID, nil
				} else if transactionStatusResp.FraudStatus == "accept" {
					fmt.Println("Payment received")
					paymentStatus = 2
					return paymentStatus, transactionStatusResp.OrderID, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				fmt.Println("Payment status settlement")
				paymentStatus = 2
				return paymentStatus, transactionStatusResp.OrderID, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				fmt.Println("Payment status denied")
				paymentStatus = 3
				return paymentStatus, transactionStatusResp.OrderID, nil
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				fmt.Println("Payment status failure")
				paymentStatus = 4
				return paymentStatus, transactionStatusResp.OrderID, nil
			} else if transactionStatusResp.TransactionStatus == "pending" {
				fmt.Println("Payment status pending")
				paymentStatus = 5
				return paymentStatus, transactionStatusResp.OrderID, nil
			}
		}
	}

	return 0, "", nil
}
