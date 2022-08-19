package usecase

import (
	"context"
	"database/sql"
	"simple-transaction-api/repository"
	"simple-transaction-api/repository/models"
	"simple-transaction-api/usecase/requests"

	"github.com/sirupsen/logrus"
)

// BrandUC ...
type BrandUC struct {
	*ContractUC
	*sql.Tx
}

// BuildBody ...
func (uc BrandUC) BuildBody(res *models.Product) {}

// AddBrand ...
func (uc BrandUC) AddBrand(c context.Context, input *requests.BrandRequest) (res models.Brand, err error) {
	repo := repository.NewBrandRepository(uc.DB, uc.Tx)

	res = models.Brand{
		Name:        input.Name,
		Description: input.Description,
		MadeIn:      input.MadeIn,
		Status:      true,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logrus.Println("[Add.AddBrand] Err : ", err)
		return res, err
	}

	return res, err
}
