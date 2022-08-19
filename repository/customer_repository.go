package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository/models"
)

// ICustomerRepository ...
type ICustomerRepository interface {
	Add(c context.Context, model *models.Customer) (res int, err error)
	FindByID(c context.Context, id int) (res models.Customer, err error)
	FindAll(c context.Context, param models.CustomerParameter) (res []models.Customer, err error)
}

// CustomerRepository ...
type CustomerRepository struct {
	DB *sql.DB
	Tx *sql.Tx
}

// NewCustomerRepository ...
func NewCustomerRepository(DB *sql.DB, Tx *sql.Tx) ICustomerRepository {
	return &CustomerRepository{DB: DB, Tx: Tx}
}

func (repository CustomerRepository) scanRow(row *sql.Row) (res models.Customer, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.Username, &res.Password, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository CustomerRepository) scanRows(rows *sql.Rows) (res models.Customer, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.Username, &res.Password, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

// Add ...
func (repository CustomerRepository) Add(c context.Context, model *models.Customer) (res int, err error) {

	statement := `INSERT INTO customers (name, username, password, status) VALUES ($1, $2, $3, $4) returning id`

	if repository.Tx != nil {
		err = repository.Tx.QueryRowContext(c, statement,
			model.Name, model.Username, model.Password, model.Status,
		).Scan(&res)
	} else {
		err = repository.DB.QueryRowContext(c, statement,
			model.Name, model.Username, model.Password, model.Status,
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}

// FindByID ...
func (repository CustomerRepository) FindByID(c context.Context, id int) (res models.Customer, err error) {

	statement := str.Spacing(models.CustomerSelectStatement, models.CustomerWhereStatement, ` AND id = $1`)
	row := repository.DB.QueryRowContext(c, statement, id)
	res, err = repository.scanRow(row)
	if err != nil {
		return res, err
	}

	return res, err
}

// FindAll ...
func (repository CustomerRepository) FindAll(c context.Context, param models.CustomerParameter) (res []models.Customer, err error) {
	var condition string
	if param.Search != "" {
		condition += ` AND name ilike '%` + param.Search + `%'`
	}

	statement := str.Spacing(models.CustomerSelectStatement, models.CustomerWhereStatement, condition)

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
