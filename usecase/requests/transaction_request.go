package requests

// TransactionRequest ...
type TransactionRequest struct {
	ID                int     `json:"id"`
	TransactionNumber string  `json:"transaction_number"`
	CustomerID        int     `json:"customer_id"`
	PicName           string  `json:"pic_name"`
	PicPhone          string  `json:"pic_phone"`
	PicEmail          string  `json:"pic_email"`
	TotalPrice        float64 `json:"total_price"`
	TypeOfPayment     string  `json:"type_of_payment"`
	Note              string  `json:"note"`
	Status            bool    `json:"status"`
}
