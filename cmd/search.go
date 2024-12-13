package cmd

import (
	"fmt"
	"strings"
	"todo-cli/internal/models"

	"todo-cli/internal/storage"
	"todo-cli/internal/utils"
)

// TODO change it to use database
func Search(query string) {
	todos, _ := storage.ReadTodosFromJSON()
	var matchedTodos []models.Todo
	for _, todo := range todos {
		if strings.Contains(strings.ToLower(todo.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(todo.Description), strings.ToLower(query)) {
			matchedTodos = append(matchedTodos, todo)
		}
	}

	if len(matchedTodos) == 0 {
		fmt.Println("No todos found matching the query.")
		return
	}

	utils.DisplayTodos(matchedTodos)
}
