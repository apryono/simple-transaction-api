package viewmodel

// ProductVM ...
type ProductVM struct {
	ID                  int     `json:"id"`
	BrandID             int     `json:"brand_id"`
	Name                string  `json:"name"`
	OverviewDescription string  `json:"overview_description"`
	Price               float64 `json:"price"`
	Sku                 string  `json:"sku"`
}
