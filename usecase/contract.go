package usecase

import "database/sql"

// ContractUC ...
type ContractUC struct {
	EnvConfig map[string]string
	DB        *sql.DB
}
