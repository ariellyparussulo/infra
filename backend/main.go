package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Name string
	Id   int
}

func HandleUsersGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("recebendo GET request")

	params := mux.Vars(r)
	res := map[string]string{"ok": params["id"]}
	json.NewEncoder(w).Encode(res)
}

func HandleUsersPostRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("recebendo POST request")

	var post Post
	error := json.NewDecoder(r.Body).Decode(&post)

	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return

	}

	log.Printf("post", post)

	json.NewEncoder(w).Encode(&post)
}

func main() {
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
