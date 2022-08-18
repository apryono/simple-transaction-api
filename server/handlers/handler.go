package handlers

import (
	"database/sql"

	"simple-transaction-api/usecase"

	"github.com/gorilla/mux"
)

type Handler struct {
	MuxApp     *mux.Router
	ContractUC *usecase.ContractUC
	Db         *sql.DB
}
