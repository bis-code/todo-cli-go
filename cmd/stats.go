package cmd

import (
	"fmt"

	"todo-cli/internal/storage"
)

func Stats() {
	todos, _ := storage.ReadTodosFromJSON()
	total := len(todos)
	if total == 0 {
		fmt.Println("No todos available to display statistics.")
		return
	}

	completed := 0
	pending := 0
	categoryCounts := make(map[string]int)

	for _, todo := range todos {
		if todo.Completed {
			completed++
		} else {
			pending++
		}
		categoryCounts[todo.Category]++
	}

	completedPercentage := (float64(completed) / float64(total)) * 100
	pendingPercentage := (float64(pending) / float64(total)) * 100

	var mostCommonCategory string
	maxCount := 0
	for category, count := range categoryCounts {
		if count > maxCount {
			mostCommonCategory = category
			maxCount = count
		}
	}

	// Display stats
	fmt.Println("Todo Statistics:")
	fmt.Printf("Total Todos: %d\n", total)
	fmt.Printf("Completed: %d (%.2f%%)\n", completed, completedPercentage)
	fmt.Printf("Pending: %d (%.2f%%)\n", pending, pendingPercentage)

	fmt.Println("\nTodos by Category:")
	for category, count := range categoryCounts {
		fmt.Printf("%s: %d\n", category, count)
	}

	fmt.Printf("\nMost Common Category: %s (%d todos)\n", mostCommonCategory, maxCount)
}
