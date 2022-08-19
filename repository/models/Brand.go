package models

import "simple-transaction-api/usecase/viewmodel"

// Brand ...
type Brand struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	MadeIn      string                `json:"made_in"`
	Status      bool                  `json:"status"`
	Products    []viewmodel.ProductVM `json:"products,omitempty"`
	CreatedAt   string                `json:"created_at,omitempty"`
	UpdatedAt   string                `json:"updated_at,omitempty"`
	DeletedAt   *string               `json:"deleted_at,omitempty"`
}

// BrandParameter ...
type BrandParameter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MadeIn      string `json:"made_in"`
	Status      bool   `json:"status"`
	Search      string `json:"search"`
}

var (
	BrandSelectStatement = `SELECT id, name, description, made_in, status, created_at, updated_at FROM brands`

	BrandDefSelectStatement = `SELECT 
	def.id, def.name, def.description, def.made_in, 
		jsonb_agg( distinct 
			jsonb_build_object(
				'id', pr.id,
				'brand_id', pr.brand_id,
				'name', pr.name,
				'overview_description', pr.overview_description,
				'price', pr.price,
				'sku', pr.sku
			)
		), 
		def.status, def.created_at, def.updated_at
	FROM brands def
	LEFT JOIN products pr ON pr.brand_id = def.id
`

	BrandWhereStatement = `WHERE deleted_at IS NULL`

	BrandDefWhereStatement = `WHERE def.deleted_at IS NULL`

	BrandDefGroupByStatement = `GROUP BY def.id`
)
