package models

import "time"

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	Category    string
	Tags        []string
	CreatedAt   time.Time  // Timestamp when the todo was created
	CompletedAt *time.Time // Timestamp when the todo was completed (nullable)
}
