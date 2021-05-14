package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"users"
)

func main() {
	fmt.Println("Listening to port 8080...")

	http.HandleFunc("/v1/users", func(write http.ResponseWriter, read *http.Request) {
		switch read.Method {
		case "GET":
			fmt.Println("Getting users...")
			fmt.Fprint(write, string(users.ListUsers()))
		case "POST":
			fmt.Println("Insert user")
			request := json.NewDecoder(read.Body)
			newUser := users.User{}
			request.Decode(&newUser)
			fmt.Print(newUser)

			response := users.InsertUser(newUser)
			fmt.Fprint(write, string(response))
		default:
			fmt.Fprint(write, "Sorry, only GET and POST methods are supported.")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
