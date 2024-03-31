package router

import (
	"go-green/controllers"
	"go-green/database"

	"github.com/gorilla/mux"
)

func UserRouteHandler(router *mux.Router, pgCtx *database.PgCtx) {
	controllerCtx := &controllers.Context{DB: pgCtx.DB}

	router.HandleFunc("", controllerCtx.CreateUserHandler).Methods("POST")
	router.HandleFunc("/{id}", controllerCtx.GetUserHandler).Methods("GET")
	router.HandleFunc("/{id}", controllerCtx.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/{id}", controllerCtx.DeleteUserHandler).Methods("DELETE")
}
