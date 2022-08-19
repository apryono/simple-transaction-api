package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/repository/models"
)

// ITransactionDetailRepository ...
type ITransactionDetailRepository interface {
	Add(c context.Context, model *models.TransactionDetail) (res int, err error)
}

// TransactionDetailRepository ...
type TransactionDetailRepository struct {
	DB *sql.DB
	Tx *sql.Tx
}

// NewTransactionDetailRepository ...
func NewTransactionDetailRepository(DB *sql.DB, Tx *sql.Tx) ITransactionDetailRepository {
	return &TransactionDetailRepository{DB: DB, Tx: Tx}
}

// Add ...
func (repository TransactionDetailRepository) Add(c context.Context, model *models.TransactionDetail) (res int, err error) {

	statement := `INSERT INTO transaction_details (transaction_id, product_id, product_name, product_price, product_quantity
		) VALUES ($1, $2, $3, $4, $5) returning id`

	if repository.Tx != nil {
		err = repository.Tx.QueryRowContext(c, statement,
			model.TransactionID, model.ProductID, model.ProductName, model.ProductPrice, model.ProductQuantity,
		).Scan(&res)
	} else {
		err = repository.DB.QueryRowContext(c, statement,
			model.TransactionID, model.ProductID, model.ProductName, model.ProductPrice, model.ProductQuantity,
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
