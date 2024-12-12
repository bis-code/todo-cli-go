package cmd

import (
	"fmt"
	"time"

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
	dailyCounts := make(map[string]int)
	var totalCompletionTime time.Duration

	for _, todo := range todos {
		// Count completed and pending todos
		if todo.Completed {
			completed++
			// For simplicity, assume todos have a CreatedAt field added
			if todo.CreatedAt != (time.Time{}) && todo.CompletedAt != nil {
				totalCompletionTime += todo.CompletedAt.Sub(todo.CreatedAt)
			}
		} else {
			pending++
		}

		// Count todos by category
		categoryCounts[todo.Category]++

		// Count todos by creation date
		createdDay := todo.CreatedAt.Format("2006-01-02")
		dailyCounts[createdDay]++
	}

	// Calculate average completion time
	averageCompletionTime := time.Duration(0)
	if completed > 0 {
		averageCompletionTime = totalCompletionTime / time.Duration(completed)
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
	fmt.Printf("Average Completion Time: %s\n", averageCompletionTime)

	fmt.Println("\nTodos by Category:")
	for category, count := range categoryCounts {
		fmt.Printf("%s: %d\n", category, count)
	}

	fmt.Println("\nDaily Todo Creation:")
	for day, count := range dailyCounts {
		fmt.Printf("%s: %d\n", day, count)
	}

	fmt.Printf("\nMost Common Category: %s (%d todos)\n", mostCommonCategory, maxCount)
}
