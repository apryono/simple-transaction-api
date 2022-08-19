package usecase

import (
	"context"
	"database/sql"
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
		logrus.Println("[Add.AddProduct] Err : ", err)
		return res, err
	}

	return res, err
}
