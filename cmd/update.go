package cmd

import (
	"fmt"
	"time"

	"todo-cli/internal/storage"
)

// TODO change it to use database
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

			if todo.Completed != completed && completed {
				now := time.Now()
				todos[i].CompletedAt = &now // Set completion timestamp
			}

			todos[i].Completed = completed
			storage.WriteTodosToJSON(todos)
			fmt.Println("Todo updated successfully!")
			return
		}
	}
	fmt.Println("Todo not found.")
}
