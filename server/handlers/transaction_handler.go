package handlers

import (
	"context"
	"net/http"
	"simple-transaction-api/helper"
	"simple-transaction-api/pkg/str"
	"simple-transaction-api/repository/models"
	"simple-transaction-api/usecase"

	"github.com/gorilla/mux"
)

// TransactionHandler ...
type TransactionHandler struct {
	Handler
}

// FindByID ...
func (h *TransactionHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	id := str.StringToInt(mux.Vars(r)["transaction_id"])
	if id == 0 {
		helper.ResponseErr(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}

	uc := usecase.TransactionUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, models.TransactionParameter{ID: id})
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindAllTransaction ...
func (h *TransactionHandler) FindAllTransaction(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	search, ok := r.URL.Query()["search"]
	if !ok || len(search[0]) < 1 {
		search = append(search, "")
	}
	customerID, ok := r.URL.Query()["customer_id"]
	if !ok || len(customerID[0]) < 1 {
		customerID = append(customerID, "0")
	}

	param := models.TransactionParameter{
		Search:     search[0],
		CustomerID: str.StringToInt(customerID[0]),
	}

	uc := usecase.TransactionUC{ContractUC: h.ContractUC}
	res, err := uc.FindAllTransaction(c, param)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}
