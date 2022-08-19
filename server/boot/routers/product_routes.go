package routers

import (
	"net/http"
	"os"
	hd "simple-transaction-api/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ProductRoutes ...
type ProductRoutes struct {
	RouterGroup *mux.Router
	Handler     hd.Handler
}

// RegisterRoute ...
func (route ProductRoutes) RegisterRoute() {
	handler := hd.ProductHandler{Handler: route.Handler}

	v1 := route.RouterGroup.PathPrefix("/api/product").Subrouter()
	r := v1.NewRoute().Subrouter()
	r.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.AddProduct))).Methods(http.MethodPost)
}
