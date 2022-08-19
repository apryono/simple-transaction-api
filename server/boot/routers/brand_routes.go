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
	r.Handle("/id/{brand_id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.FindByID))).Methods(http.MethodGet)
	r.Handle("/findAll", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.FindAllBrand))).Methods(http.MethodGet)
	r.Handle("/product/id/{brand_id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.FindByIDWithProduct))).Methods(http.MethodGet)
}
