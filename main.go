package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/users", getUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/users", createUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	// Start the server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get all users")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "Get user with ID: %s", userID)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user")
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "Update user with ID: %s", userID)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "Delete user with ID: %s", userID)
}

