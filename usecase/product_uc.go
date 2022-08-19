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

// ProductUC ...
type ProductUC struct {
	*ContractUC
	*sql.Tx
}

// BuildBody ...
func (uc ProductUC) BuildBody(res *models.Product) {}

// AddProduct ...
func (uc ProductUC) AddProduct(c context.Context, input *requests.ProductRequest) (res models.Product, err error) {
	repo := repository.NewProductRepository(uc.DB, uc.Tx)

	res = models.Product{
		BrandID:             input.BrandID,
		Name:                input.Name,
		OverviewDescription: input.OverviewDescription,
		Price:               input.Price,
		Sku:                 input.Sku,
		Status:              true,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logrus.Println("[Add.AddProduct.ProductUC] Err : ", err)
		return res, err
	}

	return res, err
}

// FindByID ...
func (uc ProductUC) FindByID(c context.Context, data models.ProductParameter) (res models.Product, err error) {
	repo := repository.NewProductRepository(uc.DB, uc.Tx)
	res, err = repo.FindByID(c, data.ID)
	if err != nil {
		logrus.Println("[FindByID.FindByID.ProductUC] Err : ", err)
		if err == sql.ErrNoRows {
			return res, errors.New("Data not found")
		}
		return res, errors.New("Something went error")
	}

	return res, err
}

// FindAllProduct ...
func (uc ProductUC) FindAllProduct(c context.Context, param models.ProductParameter) (res []models.Product, err error) {
	repo := repository.NewProductRepository(uc.DB, uc.Tx)
	res, err = repo.FindAll(c, param)
	if err != nil {
		logrus.Println("[FindAll.FindAllProduct.ProductUC] Err : ", err)
		return res, err
	}

	if len(res) < 1 {
		return res, errors.New("Data not found")
	}

	for i := range res {
		uc.BuildBody(&res[i])
	}

	return res, err
}
