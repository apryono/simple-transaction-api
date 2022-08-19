package usecase

import (
	"context"
	"database/sql"
	"errors"
	"simple-transaction-api/repository/models"
	"simple-transaction-api/usecase/requests"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// CheckoutUC ...
type CheckoutUC struct {
	*ContractUC
	*sql.Tx
}

func (uc CheckoutUC) checkDetail(c context.Context, input *requests.CheckoutRequest) (err error) {
	var total float64
	now := time.Now() // current local time
	orderID := strconv.FormatInt(now.Unix(), 10)
	input.TransactionNumber = "TRSC-" + orderID

	for _, detail := range input.TransactionDetails {
		price := detail.ProductPrice
		quantity := detail.ProductQuantity
		times := (price * float64(quantity))

		total += times
	}

	input.TotalPrice = total

	return err
}

// Checkout ...
func (uc CheckoutUC) Checkout(c context.Context, input *requests.CheckoutRequest) (res models.Transaction, err error) {
	if err = uc.checkDetail(c, input); err != nil {
		logrus.Println("[checkDetail.Checkout.CheckoutUC] Err : ", err)
		return res, errors.New("Something went error")
	}

	transactionUC := TransactionUC{ContractUC: uc.ContractUC, Tx: uc.Tx}
	res, err = transactionUC.AddTransaction(c, &input.TransactionRequest)
	if err != nil {
		logrus.Println("[AddTransaction.Checkout.CheckoutUC] Err : ", err)
		return res, errors.New("Something went error")
	}

	transactionDetailUC := TransactionDetailUC{ContractUC: uc.ContractUC, Tx: uc.Tx}
	res.TransactionDetails, err = transactionDetailUC.AddMultiple(c, res.ID, &input.TransactionDetails)
	if err != nil {
		logrus.Println("[AddMultiple.Checkout.CheckoutUC] Err : ", err)
		return res, errors.New("Something went error")
	}

	return res, err
}
