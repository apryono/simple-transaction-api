package routers

import (
	"net/http"
	"os"
	hd "simple-transaction-api/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// TransactionRoutes ...
type TransactionRoutes struct {
	RouterGroup *mux.Router
	Handler     hd.Handler
}

// RegisterRoute ...
func (route TransactionRoutes) RegisterRoute() {
	handler := hd.TransactionHandler{Handler: route.Handler}

	v1 := route.RouterGroup.PathPrefix("/api/transaction").Subrouter()
	r := v1.NewRoute().Subrouter()
	r.Handle("/id/{transaction_id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.FindByID))).Methods(http.MethodGet)
}
