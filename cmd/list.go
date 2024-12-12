package cmd

import (
	"sort"
	"strings"

	"todo-cli/internal/models"
	"todo-cli/internal/storage"
	"todo-cli/internal/utils"
)

func List(filter, sortBy, category, tagsStr string) {
	todos, _ := storage.ReadTodosFromJSON()
	tags := []string{}
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	// Apply filters
	if filter != "" {
		var filteredTodos []models.Todo
		for _, todo := range todos {
			if (filter == "completed" && todo.Completed) || (filter == "pending" && !todo.Completed) {
				filteredTodos = append(filteredTodos, todo)
			}
		}
		todos = filteredTodos
	}

	if category != "" && category != "all" {
		var categorizedTodos []models.Todo
		for _, todo := range todos {
			if strings.EqualFold(todo.Category, category) {
				categorizedTodos = append(categorizedTodos, todo)
			}
		}
		todos = categorizedTodos
	}

	if len(tags) > 0 && !(len(tags) == 1 && tags[0] == "") {
		var taggedTodos []models.Todo
		for _, todo := range todos {
			for _, tag := range tags {
				if utils.Contains(todo.Tags, tag) {
					taggedTodos = append(taggedTodos, todo)
					break
				}
			}
		}
		todos = taggedTodos
	}

	// Apply sorting
	switch sortBy {
	case "title":
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].Title < todos[j].Title
		})
	case "id":
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].ID < todos[j].ID
		})
	}

	// Display todos
	utils.DisplayTodos(todos)
}
