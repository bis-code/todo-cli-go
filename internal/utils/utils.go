package utils

import (
	"fmt"
	"strings"

	"todo-cli/internal/models"
)

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}

func DisplayTodos(todos []models.Todo) {
	if len(todos) == 0 {
		fmt.Println("No todos to display.")
		return
	}
	for _, todo := range todos {
		tags := strings.Join(todo.Tags, ", ")
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %v, Category: %s, Tags: [%s]\n",
			todo.ID, todo.Title, todo.Description, todo.Completed, todo.Category, tags)
	}
}
