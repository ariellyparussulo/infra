package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitTables()

	router := mux.NewRouter()
	router.HandleFunc("/v1/users/", HandleUsersPostRequest).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/v1/users/{id:[0-9]+}", HandleUsersGetRequest).
		Methods("GET").
		Schemes("http")

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",
	}

	log.Fatal(srv.ListenAndServe())
}
