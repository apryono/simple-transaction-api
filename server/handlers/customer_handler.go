package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-transaction-api/helper"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository/models"
	"simple-transaction-api/usecase"
	"simple-transaction-api/usecase/requests"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	Handler
}

// AddCustomer ...
func (h *CustomerHandler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	input := new(requests.CustomerRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.AddCustomer(c, input)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindByID ...
func (h *CustomerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	id := str.StringToInt(mux.Vars(r)["customer_id"])
	if id == 0 {
		helper.ResponseErr(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}

	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, models.CustomerParameter{ID: id}, false)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindAllCustomer ...
func (h *CustomerHandler) FindAllCustomer(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	search, ok := r.URL.Query()["search"]
	if !ok || len(search[0]) < 1 {
		search = append(search, "")
	}

	param := models.CustomerParameter{
		Search: search[0],
	}

	uc := usecase.CustomerUC{ContractUC: h.ContractUC}
	res, err := uc.FindAllCustomer(c, param)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}
