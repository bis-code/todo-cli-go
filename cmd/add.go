package cmd

import (
	"fmt"
	"time"

	"todo-cli/internal/models"
	"todo-cli/internal/storage"
)

func Add(title, desc, category, tagsStr string) {
	var tags []string
	if tagsStr != "" {
		tags = append(tags, tagsStr)
	}
	todo := models.Todo{
		Title:       title,
		Description: desc,
		Completed:   false,
		Category:    category,
		Tags:        tags,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	err := storage.AddTodo(todo)
	if err != nil {
		fmt.Println("Error adding todo:", err)
		return
	}
	fmt.Println("Todo added successfully!")
}
