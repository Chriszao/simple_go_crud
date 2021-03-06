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

	// ReadAll
	router.HandleFunc("/users", routes.FetchUsers).Methods(http.MethodGet)

	// ReadById
	router.HandleFunc("/users/{id}", routes.GetUserById).Methods(http.MethodGet)

	// Update
	router.HandleFunc("/users/{id}", routes.UpdateUser).Methods(http.MethodPut)

	// Delete
	router.HandleFunc("/users/{id}", routes.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("✨ Listening on http://localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", router))
}
