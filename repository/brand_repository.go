package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/repository/models"
)

// IBrandRepository ...
type IBrandRepository interface {
	Add(c context.Context, model *models.Brand) (res int, err error)
}

// BrandRepository ...
type BrandRepository struct {
	DB *sql.DB
	Tx *sql.Tx
}

// NewBrandRepository ...
func NewBrandRepository(DB *sql.DB, Tx *sql.Tx) IBrandRepository {
	return &BrandRepository{DB: DB, Tx: Tx}
}

// Add ...
func (repository BrandRepository) Add(c context.Context, model *models.Brand) (res int, err error) {

	statement := `INSERT INTO brands (name, description, made_in, status) VALUES ($1, $2, $3, $4) returning id`

	if repository.Tx != nil {
		err = repository.Tx.QueryRowContext(c, statement,
			model.Name, model.Description, model.MadeIn, model.Status,
		).Scan(&res)
	} else {
		err = repository.DB.QueryRowContext(c, statement,
			model.Name, model.Description, model.MadeIn, model.Status,
		).Scan(&res)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
