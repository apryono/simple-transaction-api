package requests

import "errors"

// CustomerRequest ...
type CustomerRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

// Validate ...
func (r *CustomerRequest) Validate(req *CustomerRequest) (err error) {
	if req.Name == "" {
		return errors.New("Required Name")
	}
	if req.Username == "" {
		return errors.New("Required Username")
	}
	if req.Password == "" {
		return errors.New("Required Password")
	}

	return err
}
