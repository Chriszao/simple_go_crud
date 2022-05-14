package main

import (
	"crud/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Create
	router.HandleFunc("/users", routes.InsertUser).Methods(http.MethodPost)

	fmt.Println("✨ Listening on http://localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", router))
}
