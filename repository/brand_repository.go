package repository

import (
	"context"
	"database/sql"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository/models"
)

// IBrandRepository ...
type IBrandRepository interface {
	Add(c context.Context, model *models.Brand) (res int, err error)
	FindByID(c context.Context, id int) (res models.Brand, err error)
	FindAll(c context.Context, param models.BrandParameter) (res []models.Brand, err error)
	FindByIDWithProduct(c context.Context, id int) (res models.Brand, err error)
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

func (repository BrandRepository) scanRow(row *sql.Row) (res models.Brand, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.Description, &res.MadeIn, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BrandRepository) scanRowWithProduct(row *sql.Row) (res models.Brand, err error) {
	err = row.Scan(
		&res.ID, &res.Name, &res.Description, &res.MadeIn, &models.UnmarshalModel{To: &res.Products}, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository BrandRepository) scanRows(rows *sql.Rows) (res models.Brand, err error) {
	err = rows.Scan(
		&res.ID, &res.Name, &res.Description, &res.MadeIn, &res.Status, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, err
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

// FindByID ...
func (repository BrandRepository) FindByID(c context.Context, id int) (res models.Brand, err error) {

	statement := str.Spacing(models.BrandSelectStatement, models.BrandWhereStatement, ` AND id = $1`)
	row := repository.DB.QueryRowContext(c, statement, id)
	res, err = repository.scanRow(row)
	if err != nil {
		return res, err
	}
	return res, err
}

// FindAll ...
func (repository BrandRepository) FindAll(c context.Context, param models.BrandParameter) (res []models.Brand, err error) {
	var condition string
	if param.Search != "" {
		condition += ` AND name ilike '%` + param.Search + `%'`
	}

	statement := str.Spacing(models.BrandSelectStatement, models.BrandWhereStatement, condition)

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

// FindByIDWithProduct ...
func (repository BrandRepository) FindByIDWithProduct(c context.Context, id int) (res models.Brand, err error) {

	statement := str.Spacing(models.BrandDefSelectStatement, models.BrandDefWhereStatement, ` AND def.id = $1`, models.BrandDefGroupByStatement)
	row := repository.DB.QueryRowContext(c, statement, id)
	res, err = repository.scanRowWithProduct(row)
	if err != nil {
		return res, err
	}
	return res, err
}
