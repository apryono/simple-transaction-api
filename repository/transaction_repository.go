package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/repository/models"
)

// ITransactionRepository ...
type ITransactionRepository interface {
	Add(c context.Context, model *models.Transaction) (res int, err error)
}

// TransactionRepository ...
type TransactionRepository struct {
	DB *sql.DB
	Tx *sql.Tx
}

// NewTransactionRepository ...
func NewTransactionRepository(DB *sql.DB, Tx *sql.Tx) ITransactionRepository {
	return &TransactionRepository{DB: DB, Tx: Tx}
}

// Add ...
func (repository TransactionRepository) Add(c context.Context, model *models.Transaction) (res int, err error) {

	statement := `INSERT INTO transactions (transaction_number, customer_id, pic_name, pic_phone, pic_email, total_price, type_of_payment, note, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	if repository.Tx != nil {
		err = repository.Tx.QueryRowContext(c, statement,
			model.TransactionNumber, model.CustomerID, model.PicName, model.PicPhone, model.PicEmail, model.TotalPrice,
			model.TypeOfPayment, model.Note, model.Status,
		).Scan(&res)
	} else {
		err = repository.DB.QueryRowContext(c, statement,
			model.TransactionNumber, model.CustomerID, model.PicName, model.PicPhone, model.PicEmail, model.TotalPrice,
			model.TypeOfPayment, model.Note, model.Status,
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
