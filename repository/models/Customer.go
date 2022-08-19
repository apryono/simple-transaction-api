package models

// Customer ...
type Customer struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Status    bool    `json:"status"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
