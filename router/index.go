package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-green/database"

	"github.com/gorilla/mux"
)

func Init(pgCtx *database.PgCtx) {
	// Create a new router
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(jsonMiddleware)
	apiRouter.HandleFunc("/", homeHandler).Methods("GET")

	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	UserRouteHandler(userRouter, pgCtx)

	// Start the server
	port := os.Getenv("PORT")
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]string{
		"message": "Welcome to the Go Green API",
	}
	json.NewEncoder(w).Encode(resp)
}
