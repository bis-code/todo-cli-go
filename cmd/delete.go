package cmd

import (
	"fmt"
	"todo-cli/internal/models"

	"todo-cli/internal/storage"
)

// TODO change it to use database
func Delete(id int, deleteAll bool) {
	if deleteAll {
		storage.WriteTodosToJSON([]models.Todo{})
		fmt.Println("All todos deleted successfully!")
		return
	}

	todos, _ := storage.ReadTodosFromJSON()
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			storage.WriteTodosToJSON(todos)
			fmt.Println("Todo deleted successfully!")
			return
		}
	}
	fmt.Println("Todo not found.")
}
