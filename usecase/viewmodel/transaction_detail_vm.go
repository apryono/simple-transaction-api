package viewmodel

// TransactionDetailVM ...
type TransactionDetailVM struct {
	ID              int     `json:"id"`
	TransactionID   int     `json:"transaction_id"`
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductPrice    float64 `json:"product_price"`
	ProductQuantity int     `json:"product_quantity"`
}
