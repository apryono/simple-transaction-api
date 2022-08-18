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

	apiV1 := boot.App.PathPrefix("/v1")

	custRoutes := routers.CustomerRoutes{RouterGroup: apiV1, Handler: handler}
	custRoutes.RegisterRoute()

}
