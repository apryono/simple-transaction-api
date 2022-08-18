package boot

import (
	"simple-transaction-api/usecase"

	"github.com/gorilla/mux"
)

// Boot ...
type Boot struct {
	App        *mux.Router
	ContractUC usecase.ContractUC
}
