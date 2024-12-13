package storage

import (
	"context"
	"fmt"
	"log"
	"time"
	"todo-cli/internal/models"

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

func AddTodo(todo models.Todo) error {
	_, err := conn.Exec(context.Background(), `
		INSERT iNTO todos (title, description, completed, category,tags, created_at, completed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, todo.Title, todo.Description, todo.Completed, todo.Category, todo.Tags, todo.CreatedAt, todo.CompletedAt)
	return err
}

func GetAllTodos() ([]models.Todo, error) {
	rows, err := conn.Query(context.Background(), `
		SELECT id, title, description, completed, category, tags, created_at, completed_at
		FROM todos
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var completedAt *time.Time
		var tags []string

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Category, &todo.Tags, &tags, &todo.CreatedAt, &completedAt)
		if err != nil {
			return nil, err
		}
		todo.Tags = tags
		todo.CompletedAt = completedAt
		todos = append(todos, todo)
	}

	return todos, nil
}

func UpdateTodo(todo models.Todo) error {
	_, err := conn.Exec(context.Background(), `
		UPDATE todos
		SET title = $1, description = $2, completed = $3, category = $4, tags = $5, completed_at = $6
		WHERE id = $7
	`, todo.Title, todo.Description, todo.Completed, todo.Category, todo.Tags, todo.CompletedAt, todo.ID)
	return err
}

func DeleteTodoByID(id int) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM todos WHERE id = $1", id)
	return err
}

//func GetTodoStats() (map[string]interface{}, error) {
//	query := `
//		SELECT
//			COUNT(*) AS total,
//			SUM(CASE WHEN completed THEN 1 ELSE 0 END) as completed,
//			SUM(CASE WHEN NOT completed THEN 1 ELSE 0 END) AS pending
//		FROM todos
//	`
//
//	var stats map[string]interface{}
//	err := conn.QueryRow(context.Background(), query).Scan(&stats["total"], &stats["completed"], &stats["pending"])
//	return stats, err
//}
