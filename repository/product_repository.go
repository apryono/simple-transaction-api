package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository/models"
	"strconv"
)

// IProductRepository ...
type IProductRepository interface {
	Add(c context.Context, model *models.Product) (res int, err error)
	FindByID(c context.Context, id int) (res models.Product, err error)
	FindAll(c context.Context, param models.ProductParameter) (res []models.Product, err error)
}

// ProductRepository ...
type ProductRepository struct {
	DB *sql.DB
	Tx *sql.Tx
}

// NewProductRepository ...
func NewProductRepository(DB *sql.DB, Tx *sql.Tx) IProductRepository {
	return &ProductRepository{DB: DB, Tx: Tx}
}

func (repository ProductRepository) scanRow(row *sql.Row) (res models.Product, err error) {
	err = row.Scan(
		&res.ID, &res.BrandID, &models.UnmarshalModel{To: &res.Brand}, &res.Name, &res.OverviewDescription, &res.Price,
		&res.Sku, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository ProductRepository) scanRows(rows *sql.Rows) (res models.Product, err error) {
	err = rows.Scan(
		&res.ID, &res.BrandID, &models.UnmarshalModel{To: &res.Brand}, &res.Name, &res.OverviewDescription, &res.Price,
		&res.Sku, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

// Add ...
func (repository ProductRepository) Add(c context.Context, model *models.Product) (res int, err error) {

	statement := `INSERT INTO products (name, brand_id, overview_description, price, sku, status) VALUES ($1, $2, $3, $4, $5, $6) returning id`

	if repository.Tx != nil {
		err = repository.Tx.QueryRowContext(c, statement,
			model.Name, model.BrandID, model.OverviewDescription, model.Price, model.Sku, model.Status,
		).Scan(&res)
	} else {
		err = repository.DB.QueryRowContext(c, statement,
			model.Name, model.BrandID, model.OverviewDescription, model.Price, model.Sku, model.Status,
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}

// FindByID ...
func (repository ProductRepository) FindByID(c context.Context, id int) (res models.Product, err error) {

	statement := str.Spacing(models.ProductSelectStatement, models.ProductWhereStatement, ` AND def.id = $1`, models.ProductGroupByStatement)
	row := repository.DB.QueryRowContext(c, statement, id)
	res, err = repository.scanRow(row)
	if err != nil {
		return res, err
	}

	return res, err
}

// FindAll ...
func (repository ProductRepository) FindAll(c context.Context, param models.ProductParameter) (res []models.Product, err error) {
	var condition string
	if param.Search != "" {
		condition += ` AND def.name ilike '%` + param.Search + `%'`
	}

	if param.BrandID != 0 {
		condition += ` AND br.id = ` + strconv.Itoa(param.BrandID)
	}

	statement := str.Spacing(models.ProductSelectStatement, models.ProductWhereStatement, condition, models.ProductGroupByStatement)

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
