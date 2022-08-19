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
