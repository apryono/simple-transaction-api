package routers

import (
	"net/http"
	"os"
	hd "simple-transaction-api/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CustomerRoutes ...
type CustomerRoutes struct {
	RouterGroup *mux.Router
	Handler     hd.Handler
}

// RegisterRoute ...
func (route CustomerRoutes) RegisterRoute() {
	handler := hd.CustomerHandler{Handler: route.Handler}

	v1 := route.RouterGroup.PathPrefix("/api/customer").Subrouter()
	r := v1.NewRoute().Subrouter()
	r.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.AddCustomer))).Methods(http.MethodPost)
	r.Handle("/id/{customer_id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.FindByID))).Methods(http.MethodGet)
	r.Handle("/findAll", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.FindAllCustomer))).Methods(http.MethodGet)
}
