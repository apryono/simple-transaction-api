package routers

import (
	"simple-transaction-api/server/handlers"

	"github.com/gorilla/mux"
)

// CustomerRoutes ...
type CustomerRoutes struct {
	RouterGroup *mux.Route
	Handler     handlers.Handler
}

// RegisterRoute ...
func (route CustomerRoutes) RegisterRoute() {

	v1 := route.RouterGroup.PathPrefix("/user").Subrouter()
	_ = v1.NewRoute().Subrouter()
}
