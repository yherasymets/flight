package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *app) routes() *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(a.notFoundResp)
	router.MethodNotAllowedHandler = http.HandlerFunc(a.methodNotAllowedResp)

	router.HandleFunc("/v1/info", a.infoHandler).Methods("GET")
	router.HandleFunc("/v1/flight", a.craeteFlight).Methods("POST")
	router.HandleFunc("/v1/flight/{id}", a.getFlihgt).Methods("GET")
	router.HandleFunc("/v1/flight/{id}", a.updateFlight).Methods("PUT")
	router.HandleFunc("/v1/flight/{id}", a.deleteFlihgt).Methods("DELETE")
	return router
}
