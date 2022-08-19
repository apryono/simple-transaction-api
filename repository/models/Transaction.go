package models

// Transaction ...
type Transaction struct {
	ID                 int                 `json:"id"`
	TransactionNumber  string              `json:"transaction_number"`
	CustomerID         int                 `json:"customer_id"`
	Customer           Customer            `json:"customer"`
	PicName            string              `json:"pic_name"`
	PicPhone           string              `json:"pic_phone"`
	PicEmail           string              `json:"pic_email"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	TotalPrice         float64             `json:"total_price"`
	TypeOfPayment      string              `json:"type_of_payment"`
	Note               string              `json:"note"`
	Status             string              `json:"status"`
	CreatedAt          string              `json:"created_at"`
	UpdatedAt          string              `json:"updated_at"`
	DeletedAt          *string             `json:"deleted_at"`
}

// TransactionParameter ...
type TransactionParameter struct {
	ID                 int                 `json:"id"`
	TransactionNumber  string              `json:"transaction_number"`
	CustomerID         int                 `json:"customer_id"`
	Customer           Customer            `json:"customer"`
	PicName            string              `json:"pic_name"`
	PicPhone           string              `json:"pic_phone"`
	PicEmail           string              `json:"pic_email"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	TotalPrice         float64             `json:"total_price"`
	TypeOfPayment      string              `json:"type_of_payment"`
	Note               string              `json:"note"`
	Status             string              `json:"status"`
	Search             string              `json:"search"`
}

var (
	StatusPending  = "pending"
	StatusPaid     = "paid"
	StatusCanceled = "canceled"
)
