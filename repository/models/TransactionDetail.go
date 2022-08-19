package models

// TransactionDetail ...
type TransactionDetail struct {
	ID              int     `json:"id"`
	TransactionID   int     `json:"transaction_id"`
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float64 `json:"product_price"`
	ProductQuantity int     `json:"product_quantity"`
	CreatedAt       string  `json:"created_at,omitempty"`
	UpdatedAt       string  `json:"updated_at,omitempty"`
	DeletedAt       *string `json:"deleted_at,omitempty"`
}

// TransactionDetailParameter ...
type TransactionDetailParameter struct {
	ID              int     `json:"id"`
	TransactionID   int     `json:"transaction_id"`
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float64 `json:"product_price"`
	ProductQuantity int     `json:"product_quantity"`
	Search          string  `json:"search"`
}
