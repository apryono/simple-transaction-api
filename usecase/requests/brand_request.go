package requests

import "errors"

// BrandRequest ...
type BrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	MadeIn      string `json:"made_in"`
	Status      bool   `json:"status"`
}

// Validate ...
func (r *BrandRequest) Validate(req *BrandRequest) (err error) {
	if req.Name == "" {
		return errors.New("Required Name")
	}
	if req.Description == "" {
		return errors.New("Required Description")
	}
	if req.MadeIn == "" {
		return errors.New("Required Made In")
	}

	return err
}
