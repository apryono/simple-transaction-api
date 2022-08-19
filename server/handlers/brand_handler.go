package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-transaction-api/helper"
	"simple-transaction-api/usecase"
	"simple-transaction-api/usecase/requests"
)

// BrandHandler ...
type BrandHandler struct {
	Handler
}

// AddBrand ...
func (h *BrandHandler) AddBrand(w http.ResponseWriter, r *http.Request) {
	c := context.Background()
	fmt.Println("masuk AddBrand")
	input := new(requests.BrandRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(input); err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.BrandUC{ContractUC: h.ContractUC}
	res, err := uc.AddBrand(c, input)
	if err != nil {
		helper.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}
