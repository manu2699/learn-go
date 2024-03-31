package main

import (
	// "log"

	"go-green/database"
	"go-green/router"
)

func main() {
	pgDB, _ := database.InitPostgres()

	// Set postgres db context
	pgCtx := &database.PgCtx{DB: pgDB}

	// Terminate db connection after the main function finishes/throws errors
	defer pgDB.Close()

	router.Init(pgCtx)
}
