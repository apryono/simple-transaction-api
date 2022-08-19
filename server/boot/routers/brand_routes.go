package routers

import (
	"net/http"
	"os"
	hd "simple-transaction-api/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// BrandRoutes ...
type BrandRoutes struct {
	RouterGroup *mux.Router
	Handler     hd.Handler
}

// RegisterRoute ...
func (route BrandRoutes) RegisterRoute() {
	handler := hd.BrandHandler{Handler: route.Handler}

	v1 := route.RouterGroup.PathPrefix("/api/brand").Subrouter()
	r := v1.NewRoute().Subrouter()
	r.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.AddBrand))).Methods(http.MethodPost)
}
