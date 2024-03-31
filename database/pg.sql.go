package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Postgres context
type PgCtx struct {
	DB *sql.DB
}

func InitPostgres() (db *sql.DB, err error) {
	// Connect to the PostgreSQL database
	pgDb, err := sql.Open("postgres", os.Getenv("POSTGRES_DB_URL"))

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	// Create the table if it doesn't exist
	_, tableError := pgDb.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL
			password VARCHAR(300)
		)
	`)

	if tableError != nil {
		log.Println("Failed to create table:", tableError)
		// return nil, tableError
	}

	log.Println("Database initialized successfully!")

	return pgDb, nil
}
