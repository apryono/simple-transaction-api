package models

// Customer ...
type Customer struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"password,omitempty"`
	Status    bool    `json:"status"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// CustomerParameter ...
type CustomerParameter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
	Search   string `json:"search"`
}

var (
	CustomerSelectStatement = `SELECT id, name, username, password, status, created_at, updated_at FROM customers`

	CustomerWhereStatement = `WHERE deleted_at IS NULL`
)
