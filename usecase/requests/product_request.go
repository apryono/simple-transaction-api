package requests

import "errors"

// ProductRequest ...
type ProductRequest struct {
	BrandID             int     `json:"brand_id"`
	Name                string  `json:"name"`
	OverviewDescription string  `json:"overview_description"`
	Price               float64 `json:"price"`
	Sku                 string  `json:"sku"`
	Status              bool    `json:"status"`
}

// Validate ...
func (r *ProductRequest) Validate(req *ProductRequest) (err error) {
	if req.BrandID == 0 {
		return errors.New("Required BrandID")
	}
	if req.Name == "" {
		return errors.New("Required Name")
	}
	if req.OverviewDescription == "" {
		return errors.New("Required Description")
	}
	if req.Price == 0 {
		return errors.New("Required Price")
	}
	if req.Sku == "" {
		return errors.New("Required Price")
	}

	return err
}
