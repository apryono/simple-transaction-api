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

// BrandHandler ...
type BrandHandler struct {
	Handler
}

// AddBrand ...
func (h *BrandHandler) AddBrand(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	input := new(requests.BrandRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.BrandUC{ContractUC: h.ContractUC}
	res, err := uc.AddBrand(c, input)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindByID ...
func (h *BrandHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	id := str.StringToInt(mux.Vars(r)["brand_id"])
	if id == 0 {
		helper.ResponseErr(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}

	uc := usecase.BrandUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, models.BrandParameter{ID: id})
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindAllBrand ...
func (h *BrandHandler) FindAllBrand(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	search, ok := r.URL.Query()["search"]
	if !ok || len(search[0]) < 1 {
		search = append(search, "")
	}

	param := models.BrandParameter{
		Search: search[0],
	}

	uc := usecase.BrandUC{ContractUC: h.ContractUC}
	res, err := uc.FindAllBrand(c, param)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindByIDWithProduct ...
func (h *BrandHandler) FindByIDWithProduct(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	id := str.StringToInt(mux.Vars(r)["brand_id"])
	if id == 0 {
		helper.ResponseErr(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}

	uc := usecase.BrandUC{ContractUC: h.ContractUC}
	res, err := uc.FindByIDWithProduct(c, models.BrandParameter{ID: id})
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}