package cmd

import (
	"fmt"
	"todo-cli/internal/storage"
)

func Delete(id int, deleteAll bool) {
	if deleteAll {
		err := storage.DeleteAllTodos()
		if err != nil {
			fmt.Printf("Failed to delete all todos: %v\n", err)
			return
		}
		fmt.Println("All todos deleted successfully!")
		return
	}

	err := storage.DeleteTodoByID(id)
	if err != nil {
		fmt.Printf("Failed to delete todo with ID %d: %v\n", id, err)
		return
	}
	fmt.Println("Todo deleted successfully!")
}
