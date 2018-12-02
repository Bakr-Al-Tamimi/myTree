package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grendach/myTree/myTree"
)

func main() {
	router := myTree.NewRouter() // create routes

	// these two lines are important in order to allow access from the front-end side to the methonds
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethonds := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// launch server with CORS validations
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(allowedOrigins, allowedMethonds)(router)))
}
