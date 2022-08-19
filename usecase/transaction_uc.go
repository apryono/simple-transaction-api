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

// TransactionUC ...
type TransactionUC struct {
	*ContractUC
	*sql.Tx
}

// BuildBody ...
func (uc TransactionUC) BuildBody(res *models.Transaction) {}

// AddTransaction ...
func (uc TransactionUC) AddTransaction(c context.Context, input *requests.TransactionRequest) (res models.Transaction, err error) {
	repo := repository.NewTransactionRepository(uc.DB, uc.Tx)

	res = models.Transaction{
		TransactionNumber: input.TransactionNumber,
		CustomerID:        input.CustomerID,
		PicName:           input.PicName,
		PicPhone:          input.PicPhone,
		PicEmail:          input.PicEmail,
		TotalPrice:        input.TotalPrice,
		TypeOfPayment:     input.TypeOfPayment,
		Note:              input.Note,
		Status:            models.StatusPending,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logrus.Println("[Add.AddTransaction.AddTransaction] Err : ", err)
		return res, errors.New("Something went error")
	}

	return res, err
}

// FindByID ...
func (uc TransactionUC) FindByID(c context.Context, data models.TransactionParameter) (res models.Transaction, err error) {
	repo := repository.NewTransactionRepository(uc.DB, uc.Tx)
	res, err = repo.FindByID(c, data.ID)
	if err != nil {
		logrus.Println("[FindByID.FindByID.AddTransaction] Err : ", err)
		if err == sql.ErrNoRows {
			return res, errors.New("Data not found")
		}
		return res, errors.New("Something went error")
	}

	uc.BuildBody(&res)

	return res, err
}
