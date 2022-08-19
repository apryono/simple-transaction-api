package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-transaction-api/helper"
	"simple-transaction-api/usecase"
	"simple-transaction-api/usecase/requests"
)

// CheckoutHandler ...
type CheckoutHandler struct {
	Handler
}

// AddCheckout ...
func (h *CheckoutHandler) AddCheckout(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	input := new(requests.CheckoutRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.CheckoutUC{ContractUC: h.ContractUC}
	res, err := uc.Checkout(c, input)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}
