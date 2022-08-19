package requests

import "errors"

// CheckoutRequest ...
type CheckoutRequest struct {
	TransactionRequest `json:"transaction"`
	TransactionDetails []TransactionDetailRequest `json:"transaction_details"`
}

// Validate ...
func (r *CheckoutRequest) Validate(req *CheckoutRequest) (err error) {
	if req.PicName == "" {
		return errors.New("Required PicName")
	}
	if req.CustomerID == 0 {
		return errors.New("Required CustomerID")
	}
	if req.PicPhone == "" {
		return errors.New("Required PicPhone")
	}
	if req.TypeOfPayment == "" {
		return errors.New("Required TypeOfPayment")
	}

	for _, detail := range req.TransactionDetails {
		if detail.ProductName == "" {
			return errors.New("Required ProductName")
		} else if detail.ProductID == 0 {
			return errors.New("Required ProductID")
		} else if detail.ProductPrice == 0 {
			return errors.New("Required ProductPrice")
		} else if detail.ProductQuantity == 0 {
			return errors.New("Required ProductQuantity")
		} else {
			break
		}
	}

	return err
}
