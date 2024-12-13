package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

// InitDB initializes the database connection and applies migrations
func InitDB() {
	// Connect to the database
	var err error
	conn, err = pgx.Connect(context.Background(), "postgres://ionut:1234@localhost:5432/todo_app?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	fmt.Println("Connected to the database successfully!")
}

// CloseDB closes the database connection
func CloseDB() {
	if conn != nil {
		conn.Close(context.Background())
		fmt.Println("Database connection closed.")
	}
}

// RunMigrations applies all pending migrations
func RunMigrations() error {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://ionut:1234@localhost:5432/todo_app?sslmode=disable",
	)
	if err != nil {
		return fmt.Errorf("unable to initialize migrations: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error applying migrations: %w", err)
	}
	if err == migrate.ErrNoChange {
		fmt.Println("Migrations are already up to date.")
	}
	return nil
}
