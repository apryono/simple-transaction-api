package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-transaction-api/helper"
	"simple-transaction-api/usecase"
	"simple-transaction-api/usecase/requests"
)

// ProductHandler ...
type ProductHandler struct {
	Handler
}

// AddProduct ...
func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	input := new(requests.ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(input); err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := uc.AddProduct(c, input)
	if err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}
