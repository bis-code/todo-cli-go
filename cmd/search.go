package cmd

import (
	"fmt"
	"todo-cli/internal/storage"
	"todo-cli/internal/utils"
)

// TODO solve issues - not possible to search
/**
./todo-cli search -query="Go"
Error:
Connected to the database successfully!
Failed to search todos by {3 824633935376}:
No todos to display.
Database connection closed.

*/
func Search(query string) {
	todos, err := storage.SearchTodos(query)

	if err != nil {
		fmt.Printf("Failed to search todos by %d:\n", err)
	}

	utils.DisplayTodos(todos)
}
