package cmd

import (
	"fmt"

	"todo-cli/internal/storage"
)

func Update(id int, title, desc string, completed bool) {
	todos, _ := storage.ReadTodosFromJSON()
	for i, todo := range todos {
		if todo.ID == id {
			if title != "" {
				todos[i].Title = title
			}
			if desc != "" {
				todos[i].Description = desc
			}
			todos[i].Completed = completed
			storage.WriteTodosToJSON(todos)
			fmt.Println("Todo updated successfully!")
			return
		}
	}
	fmt.Println("Todo not found.")
}
