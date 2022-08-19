package routers

import (
	"net/http"
	"os"
	hd "simple-transaction-api/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CheckoutRoutes ...
type CheckoutRoutes struct {
	RouterGroup *mux.Router
	Handler     hd.Handler
}

// RegisterRoute ...
func (route CheckoutRoutes) RegisterRoute() {
	handler := hd.CheckoutHandler{Handler: route.Handler}

	v1 := route.RouterGroup.PathPrefix("/api/checkout").Subrouter()
	r := v1.NewRoute().Subrouter()
	r.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.AddCheckout))).Methods(http.MethodPost)
}
