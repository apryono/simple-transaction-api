package models

// Brand ...
type Brand struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MadeIn      string  `json:"made_in"`
	Status      bool    `json:"status"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
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

	BrandWhereStatement = `Where deleted_at is null`
)
