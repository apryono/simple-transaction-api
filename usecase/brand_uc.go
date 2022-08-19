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

// BrandUC ...
type BrandUC struct {
	*ContractUC
	*sql.Tx
}

// BuildBody ...
func (uc BrandUC) BuildBody(res *models.Brand) {}

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
		logrus.Println("[Add.AddBrand.AddBrand] Err : ", err)
		return res, errors.New("Something went error")
	}

	return res, err
}

// FindByID ...
func (uc BrandUC) FindByID(c context.Context, data models.BrandParameter) (res models.Brand, err error) {
	repo := repository.NewBrandRepository(uc.DB, uc.Tx)
	res, err = repo.FindByID(c, data.ID)
	if err != nil {
		logrus.Println("[FindByID.FindByID.AddBrand] Err : ", err)
		if err == sql.ErrNoRows {
			return res, errors.New("Data not found")
		}
		return res, errors.New("Something went error")
	}

	uc.BuildBody(&res)

	return res, err
}

// FindAllBrand ...
func (uc BrandUC) FindAllBrand(c context.Context, param models.BrandParameter) (res []models.Brand, err error) {
	repo := repository.NewBrandRepository(uc.DB, uc.Tx)
	res, err = repo.FindAll(c, param)
	if err != nil {
		logrus.Println("[FindAll.FindAllBrand.BrandUC] Err : ", err)
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

// FindByIDWithProduct find brand with products
func (uc BrandUC) FindByIDWithProduct(c context.Context, data models.BrandParameter) (res models.Brand, err error) {
	repo := repository.NewBrandRepository(uc.DB, uc.Tx)
	res, err = repo.FindByIDWithProduct(c, data.ID)
	if err != nil {
		logrus.Println("[FindByID.FindByID.AddBrand] Err : ", err)
		if err == sql.ErrNoRows {
			return res, errors.New("Data not found")
		}
		return res, errors.New("Something went error")
	}

	uc.BuildBody(&res)

	return res, err
}
