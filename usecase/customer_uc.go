package usecase

import (
	"context"
	"database/sql"
	"errors"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository"
	"simple-transaction-api/repository/models"
	"simple-transaction-api/usecase/requests"

	"github.com/sirupsen/logrus"
)

// CustomerUC ...
type CustomerUC struct {
	*ContractUC
	*sql.Tx
}

// BuildBody ...
func (uc CustomerUC) BuildBody(res *models.Customer, showPassword bool) {
	res.Password = str.ShowString(showPassword, res.Password)
}

// AddCustomer ...
func (uc CustomerUC) AddCustomer(c context.Context, input *requests.CustomerRequest) (res models.Customer, err error) {
	repo := repository.NewCustomerRepository(uc.DB, uc.Tx)

	res = models.Customer{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Status:   true,
	}

	res.ID, err = repo.Add(c, &res)
	if err != nil {
		logrus.Println("[Add.AddCustomer.AddCustomer] Err : ", err)
		return res, errors.New("Something went error")
	}

	return res, err
}

// FindByID ...
func (uc CustomerUC) FindByID(c context.Context, data models.CustomerParameter, showPassword bool) (res models.Customer, err error) {
	repo := repository.NewCustomerRepository(uc.DB, uc.Tx)
	res, err = repo.FindByID(c, data.ID)
	if err != nil {
		logrus.Println("[FindByID.FindByID.AddCustomer] Err : ", err)
		if err == sql.ErrNoRows {
			return res, errors.New("Data not found")
		}
		return res, errors.New("Something went error")
	}

	uc.BuildBody(&res, showPassword)

	return res, err
}

// FindAllCustomer ...
func (uc CustomerUC) FindAllCustomer(c context.Context, param models.CustomerParameter) (res []models.Customer, err error) {
	repo := repository.NewCustomerRepository(uc.DB, uc.Tx)
	res, err = repo.FindAll(c, param)
	if err != nil {
		logrus.Println("[FindAll.FindAllCustomer.CustomerUC] Err : ", err)
		return res, err
	}

	if len(res) < 1 {
		return res, errors.New("Data not found")
	}

	for i := range res {
		uc.BuildBody(&res[i], false)
	}

	return res, err
}
