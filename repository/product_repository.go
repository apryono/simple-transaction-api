package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/repository/models"
)

// IProductRepository ...
type IProductRepository interface {
	Add(c context.Context, model *models.Product) (res int, err error)
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
