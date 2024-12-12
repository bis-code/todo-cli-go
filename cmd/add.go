package cmd

import (
	"fmt"
	"strings"
	"time"

	"todo-cli/internal/models"
	"todo-cli/internal/storage"
)

func Add(title, desc, category, tagsStr string) {
	todos, _ := storage.ReadTodosFromJSON()
	id := len(todos) + 1
	tags := []string{}
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}
	todo := models.Todo{
		ID:          id,
		Title:       title,
		Description: desc,
		Completed:   false,
		Category:    category,
		Tags:        tags,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	todos = append(todos, todo)
	storage.WriteTodosToJSON(todos)
	fmt.Println("Todo added successfully!")
}
