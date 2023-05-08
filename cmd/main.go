package main

import (
	"log"

	"github.com/gorilla/mux"
	config "github.com/megamsquare/go_gateway/pkg/env_config"
	"github.com/megamsquare/go_gateway/pkg/http/server"
	"github.com/megamsquare/go_gateway/routers"
)

func main() {
	err := config.Set_env(".env")
	if err != nil {
		log.Println("env file load config error: ", err)
	}

	router := mux.NewRouter()

	routers.RegisteredRouters(router)

	httpSrvConfig := server.Load_config()
	server.ListenAndServe(*httpSrvConfig, router)
}