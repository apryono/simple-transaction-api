package main

import (
	"log"
	"net/http"

	conf "simple-transaction-api/config"
	"simple-transaction-api/helper"

	"simple-transaction-api/server/boot"
	"simple-transaction-api/usecase"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// load all config
	configs, err := conf.LoadConfigs()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configs.DB.Close()

	app := mux.NewRouter().StrictSlash(true)

	ContractUC := usecase.ContractUC{
		EnvConfig: configs.EnvConfig,
		DB:        configs.DB,
	}

	// bootable
	bt := boot.Boot{
		App:        app,
		ContractUC: ContractUC,
	}

	bt.RegisterRouters()
	
	bt.App.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := helper.Message(false, "URL not found")
		helper.Response(w, http.StatusNotFound, response)
		return
	})
	
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{configs.EnvConfig["APP_CORS_DOMAIN"]})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})
	log.Println("Server start at " + configs.EnvConfig["APP_HOST"])
	log.Fatalln(http.ListenAndServe(configs.EnvConfig["APP_HOST"], handlers.CORS(headersOk, originsOk, methodsOk)(bt.App)))
}
