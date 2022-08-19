package models

// Product ...
type Product struct {
	ID                  int     `json:"id"`
	BrandID             int     `json:"brand_id"`
	Name                string  `json:"name"`
	OverviewDescription string  `json:"overview_description"`
	Price               float64 `json:"price"`
	Sku                 string  `json:"sku"`
	Status              bool    `json:"status"`
	CreatedAt           string  `json:"created_at,omitempty"`
	UpdatedAt           string  `json:"updated_at,omitempty"`
	DeletedAt           *string `json:"deleted_at,omitempty"`
}
