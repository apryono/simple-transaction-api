package boot

import (
	"net/http"
	"os"

	"simple-transaction-api/helper"
	"simple-transaction-api/server/boot/routers"
	"simple-transaction-api/server/handlers"

	hd "github.com/gorilla/handlers"
)

// RegisterRouters ...
func (boot Boot) RegisterRouters() {
	handler := handlers.Handler{
		MuxApp:     boot.App,
		ContractUC: &boot.ContractUC,
	}

	boot.App.Handle("/", hd.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := helper.Message(true, "Welcome to Simple API")
		helper.Response(w, http.StatusOK, response)
		return
	}))).Methods(http.MethodGet)

	apiV1 := boot.App.PathPrefix("/v1").Subrouter()

	// brand route
	brandRoutes := routers.BrandRoutes{RouterGroup: apiV1, Handler: handler}
	brandRoutes.RegisterRoute()

	// // product route
	productRoutes := routers.ProductRoutes{RouterGroup: apiV1, Handler: handler}
	productRoutes.RegisterRoute()

	// // Customer route
	customerRoutes := routers.CustomerRoutes{RouterGroup: apiV1, Handler: handler}
	customerRoutes.RegisterRoute()

	// // Checkout route
	checkoutRoutes := routers.CheckoutRoutes{RouterGroup: apiV1, Handler: handler}
	checkoutRoutes.RegisterRoute()

}
