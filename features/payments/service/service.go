package service

import (
	"capstone-mikti/features/payments"
	"capstone-mikti/features/tickets"
	"capstone-mikti/features/vouchers"
	"capstone-mikti/utils/midtrans"
	"errors"
	"fmt"
	"time"

	"github.com/nanorand/nanorand"
	"github.com/sirupsen/logrus"
)

type PaymentService struct {
	data     payments.PaymentDataInterface
	dataVouc vouchers.VoucherDataInterface
	dataTick tickets.TicketDataInterface
	mt       midtrans.MidtransService
}

func New(d payments.PaymentDataInterface, m midtrans.MidtransService, v vouchers.VoucherDataInterface, t tickets.TicketDataInterface) *PaymentService {
	return &PaymentService{
		data:     d,
		mt:       m,
		dataVouc: v,
		dataTick: t,
	}
}

func (p *PaymentService) CreatePayment(newData payments.Payment) (*payments.Payment, map[string]interface{}, error) {
	var paymentMaster = new(payments.PaymentMaster)
	var paymentDetail = new(payments.PaymentDetail)

	code, err := nanorand.Gen(6)
	if err != nil {
		logrus.Error("Error Generate Failed : ", err.Error())
		return nil, nil, errors.New("ERROR Generate Failed")
	}

	totalPrice := newData.Qty * newData.Price

	chargeResp, response, err := p.mt.GenerateTransaction(int(totalPrice), newData.PaymentType)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, nil, errors.New("ERROR Generate Transaction Failed")
	}

	//waktu
	layout := "2006-01-02"

	parseTransactionDate, _ := time.Parse(layout, newData.TransactionDate)

	paymentMaster.BookID = newData.BookID
	paymentMaster.UserID = newData.UserID
	paymentMaster.OrderID = "M" + code
	paymentMaster.MidtransID = chargeResp.OrderID
	paymentMaster.GrandTotal = totalPrice
	paymentMaster.TransactionDateParse = parseTransactionDate
	paymentMaster.PaymentStatus = newData.PaymentStatus
	paymentMaster.PaymentType = newData.PaymentType
	paymentMaster.Status = "Unpaid"

	paymentDetail.Total = totalPrice

	if newData.VoucherCode != "" {
		res, err := p.dataVouc.GetByCode(newData.VoucherCode)

		if err != nil {
			logrus.Error("Error Get Voucher Failed : ", err.Error())
			return nil, nil, errors.New("ERROR Get Voucher Failed")
		}

		totalPrice = (newData.Qty * newData.Price) - res.Price
		paymentMaster.GrandTotal = totalPrice

		paymentDetail.VoucherID = res.ID
		paymentDetail.VoucherPrice = res.Price
		paymentDetail.Total = totalPrice
	}

	resultMaster, err := p.data.InsertPaymentMaster(*paymentMaster)
	if err != nil {
		logrus.Error("Error Insert Payment Master : ", err.Error())
		return nil, nil, errors.New("ERROR Insert Payment Master Process Failed")
	}

	paymentID, err := p.data.GetByIDMidtrans(resultMaster.MidtransID)
	if err != nil {
		logrus.Error("Error Get By ID Midtrans : ", err.Error())
		return nil, nil, errors.New("ERROR Get By ID Midtrans Process Failed")
	}

	paymentDetail.PaymentID = uint(paymentID)
	paymentDetail.Price = newData.Price
	paymentDetail.Qty = newData.Qty

	_, err = p.data.InsertPaymentDetail(*paymentDetail)
	if err != nil {
		logrus.Error("Error Insert Payment Detail : ", err.Error())
		return nil, nil, errors.New("ERROR Insert Payment Detail Process Failed")
	}

	newData.GrandTotal = totalPrice

	return &newData, response, nil
}

func (p *PaymentService) UpdatePayment(notificationPayload map[string]interface{}, newData payments.UpdatePayment) (bool, error) {
	paymentStatus, orderId, err := p.mt.TransactionStatus(notificationPayload)
	if err != nil {
		return false, errors.New("ERROR Transaction Status Failed")
	}

	newData.PaymentStatus = uint(paymentStatus)
	result, err := p.data.GetAndUpdate(newData, orderId)
	if err != nil {
		return false, errors.New("ERROR Update Process Failed")
	}

	if newData.PaymentStatus == 2 {
		paymentId, err := p.data.GetByIDMidtrans(orderId)
		if err != nil {
			logrus.Error("Error Get By ID Midtrans : ", err.Error())
			return false, errors.New("ERROR Get By ID Midtrans Process Failed")
		}
		qtyTicket, voucherID, err := p.data.GetQtyAndVoucByID(paymentId)
		if err != nil {
			logrus.Error("Error Get Qty and Voucher By ID : ", err.Error())
			return false, errors.New("ERROR Get Qty and Voucher By ID Process Failed")
		}

		logrus.Info("Payment ID : ", paymentId)
		logrus.Info("Voucher ID : ", voucherID)
		logrus.Info("Qty Ticket : ", qtyTicket)

		ticketID, err := p.data.GetTicketByID(paymentId)
		if err != nil {
			logrus.Error("Error Get Ticket By ID : ", err.Error())
			return false, errors.New("ERROR Get Ticket By ID Process Failed")
		}

		errDecr := p.dataTick.DecreementQty(ticketID, qtyTicket)
		if errDecr != nil {
			logrus.Error("Error Decreement Quantity : ", errDecr.Error())
			return false, errors.New("ERROR Decreement Quantity Process Failed")
		}

		if voucherID != 0 {
			errUpd := p.dataVouc.UpdateQuantity(voucherID)
			if errUpd != nil {
				logrus.Error("Error Update Quantity : ", errUpd.Error())
				return false, errors.New("ERROR Update Quantity Process Failed")
			}
		}
	}

	return result, nil
}

// Admin
func (p *PaymentService) GetAll(status string, limitInt, offsetInt, timePublication int) ([]payments.PaymentInfo, error) {
	res, err := p.data.GetAll(status, limitInt, offsetInt, timePublication)

	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return nil, errors.New("ERROR Get All")
	}

	return res, nil
}

// User
func (p *PaymentService) GetByUserID(id uint, status string, limitInt, offsetInt, timePublication int) ([]payments.PaymentInfo, error) {
	res, err := p.data.GetByUserID(id, status, limitInt, offsetInt, timePublication)

	if err != nil {
		logrus.Error("Service Error : ", err.Error())
		return nil, errors.New("ERROR Get All")
	}

	return res, nil
}
