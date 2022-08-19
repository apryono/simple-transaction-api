package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository/models"
	"strconv"
)

// ITransactionRepository ...
type ITransactionRepository interface {
	Add(c context.Context, model *models.Transaction) (res int, err error)
	FindByID(c context.Context, id int) (res models.Transaction, err error)
	FindAll(c context.Context, param models.TransactionParameter) (res []models.Transaction, err error)
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

func (repository TransactionRepository) scanRow(row *sql.Row) (res models.Transaction, err error) {
	err = row.Scan(
		&res.ID, &res.TransactionNumber, &res.CustomerID, &models.UnmarshalModel{To: &res.Customer}, &res.PicName, &res.PicPhone, &res.PicEmail,
		&res.TotalPrice, &res.TypeOfPayment, &models.UnmarshalModel{To: &res.TransactionDetails},
		&res.Note, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository TransactionRepository) scanRows(rows *sql.Rows) (res models.Transaction, err error) {
	err = rows.Scan(
		&res.ID, &res.TransactionNumber, &res.CustomerID, &models.UnmarshalModel{To: &res.Customer}, &res.PicName, &res.PicPhone, &res.PicEmail,
		&res.TotalPrice, &res.TypeOfPayment, &models.UnmarshalModel{To: &res.TransactionDetails},
		&res.Note, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
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

// FindByID ...
func (repository TransactionRepository) FindByID(c context.Context, id int) (res models.Transaction, err error) {

	statement := str.Spacing(models.TransactionSelectStatement, models.TransactionWhereStatement, ` AND def.id = $1`, models.TransactionGroupByStatement)
	row := repository.DB.QueryRowContext(c, statement, id)
	res, err = repository.scanRow(row)
	if err != nil {
		return res, err
	}
	return res, err
}

// FindAll ...
func (repository TransactionRepository) FindAll(c context.Context, param models.TransactionParameter) (res []models.Transaction, err error) {
	var condition string
	if param.Search != "" {
		condition += ` AND def.transaction_number ilike '%` + param.Search + `%'`
	}

	if param.CustomerID != 0 {
		condition += ` AND def.customer_id = ` + strconv.Itoa(param.CustomerID)
	}

	statement := str.Spacing(models.TransactionSelectStatement, models.TransactionWhereStatement, condition, models.TransactionGroupByStatement)

	rows, err := repository.DB.QueryContext(c, statement)
	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, err
}
