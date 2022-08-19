package models

// Product ...
type Product struct {
	ID                  int     `json:"id"`
	BrandID             int     `json:"brand_id"`
	Brand               Brand   `json:"brand,omitempty"`
	Name                string  `json:"name"`
	OverviewDescription string  `json:"overview_description"`
	Price               float64 `json:"price"`
	Sku                 string  `json:"sku"`
	Status              bool    `json:"status"`
	CreatedAt           string  `json:"created_at,omitempty"`
	UpdatedAt           string  `json:"updated_at,omitempty"`
	DeletedAt           *string `json:"deleted_at,omitempty"`
}

// ProductParameter ...
type ProductParameter struct {
	ID                  int     `json:"id"`
	BrandID             int     `json:"brand_id"`
	Brand               Brand   `json:"brand"`
	Name                string  `json:"name"`
	OverviewDescription string  `json:"overview_description"`
	Price               float64 `json:"price"`
	Sku                 string  `json:"sku"`
	Status              bool    `json:"status"`
	Search              string  `json:"search"`
}

var (
	ProductSelectStatement = `SELECT 
		def.id, def.brand_id, jsonb_build_object(
			'id', br.id,
			'name', br.name,
			'description', br.description,
			'made_in', br.made_in,
			'status', br.status
		), def.name, def.overview_description, def.price, def.sku, def.status, def.created_at, def.updated_at  
		FROM products def
		LEFT JOIN brands br on br.id = def.brand_id 
	`
	ProductWhereStatement   = `WHERE def.deleted_at IS NULL`
	ProductGroupByStatement = `GROUP BY def.id, br.id`
)
