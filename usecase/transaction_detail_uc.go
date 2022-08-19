package usecase

import (
	"context"
	"database/sql"
	"errors"
	"simple-transaction-api/repository"
	"simple-transaction-api/repository/models"
	"simple-transaction-api/usecase/requests"

	"github.com/sirupsen/logrus"
)

// TransactionDetailUC ...
type TransactionDetailUC struct {
	*ContractUC
	*sql.Tx
}

// BuildBody ...
func (uc TransactionDetailUC) BuildBody(res *models.TransactionDetail) {}

// AddTransactionDetail ...
func (uc TransactionDetailUC) AddTransactionDetail(c context.Context, input *requests.TransactionDetailRequest) (res models.TransactionDetail, err error) {
	repo := repository.NewTransactionDetailRepository(uc.DB, uc.Tx)

	res = models.TransactionDetail{
		TransactionID:   input.TransactionID,
		ProductID:       input.ProductID,
		ProductName:     input.ProductName,
		ProductPrice:    input.ProductPrice,
		ProductQuantity: input.ProductQuantity,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logrus.Println("[Add.AddTransactionDetail.AddTransactionDetail] Err : ", err)
		return res, errors.New("Something went error")
	}

	return res, err
}

// AddMultiple ...
func (uc TransactionDetailUC) AddMultiple(c context.Context, transactionID int, inputs *[]requests.TransactionDetailRequest) (res []models.TransactionDetail, err error) {
	for _, input := range *inputs {
		input.TransactionID = transactionID
		tempRes, err := uc.AddTransactionDetail(c, &input)
		if err != nil {
			logrus.Println("[AddTransactionDetail.AddMultiple.AddTransactionDetail] Err : ", err)
			return res, errors.New("Something went error")
		}

		res = append(res, tempRes)
	}

	return res, err
}
