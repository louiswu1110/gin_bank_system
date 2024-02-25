package server

import (
	"fmt"
	"log"
	"meepshop_project/service/usecase"
	"meepshop_project/utils/config"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(usecase.MiddlewareValidateAPIKey)
	router.HandleFunc("/test-1", usecase.Test1).Methods("POST")
	port := config.GlobalConfig.GetServerPort()
	fmt.Println("server is running at port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
