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

// ProductHandler ...
type ProductHandler struct {
	Handler
}

// AddProduct ...
func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	input := new(requests.ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(input); err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	uc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := uc.AddProduct(c, input)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindByID ...
func (h *ProductHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	id := str.StringToInt(mux.Vars(r)["product_id"])
	if id == 0 {
		helper.ResponseErr(w, http.StatusBadRequest, "Invalid Parameter")
		return
	}

	uc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := uc.FindByID(c, models.ProductParameter{ID: id})
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}

// FindAllProduct ...
func (h *ProductHandler) FindAllProduct(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	search, ok := r.URL.Query()["search"]
	if !ok || len(search[0]) < 1 {
		search = append(search, "")
	}

	brandID, ok := r.URL.Query()["brand_id"]
	if !ok || len(brandID[0]) < 1 {
		brandID = append(brandID, "0")
	}

	param := models.ProductParameter{
		Search:  search[0],
		BrandID: str.StringToInt(brandID[0]),
	}

	uc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := uc.FindAllProduct(c, param)
	if err != nil {
		helper.ResponseErr(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.Response(w, http.StatusOK, res)
	return
}
