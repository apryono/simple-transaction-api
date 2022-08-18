package main

import (
	"log"
	"net/http"

	conf "simple-transaction-api/config"

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

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{configs.EnvConfig["APP_CORS_DOMAIN"]})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})
	bt.App.Use(handlers.CORS(headersOk, originsOk, methodsOk))

	bt.RegisterRouters()
	log.Println("Server start at " + configs.EnvConfig["APP_HOST"])
	log.Fatalln(http.ListenAndServe(configs.EnvConfig["APP_HOST"], bt.App))
}
