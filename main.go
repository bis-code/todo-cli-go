package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	statsCmd := flag.NewFlagSet("stats", flag.ExitOnError)

	// Add command flags
	addTitle := addCmd.String("title", "", "The title of the todo")
	addDesc := addCmd.String("desc", "", "Description of the todo")
	addCategory := addCmd.String("category", "General", "Category of the todo")
	addTags := addCmd.String("tags", "", "Comma-separated tags for the todo")

	// Update command flags
	updateID := updateCmd.Int("id", 0, "ID of the todo to update")
	updateTitle := updateCmd.String("title", "", "Updated title")
	updateDesc := updateCmd.String("desc", "", "Updated description")
	updateCompleted := updateCmd.Bool("completed", false, "Update completed task")

	// Delete command flag
	deleteID := deleteCmd.Int("id", -1, "ID of the todo to delete")
	deleteAll := deleteCmd.Bool("all", false, "Deletes all todos")

	// List command flags
	listFilter := listCmd.String("filter", "", "Filter todos by 'completed' or 'pending'")
	listSort := listCmd.String("sort", "id", "Sort todos by 'id' or 'title'")
	listCategory := listCmd.String("category", "all", "Filter todos by category")
	listTags := listCmd.String("tags", "", "Comma-separated tags to filter by")

	// Search command flag
	searchQuery := searchCmd.String("query", "", "Search query for todos")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', 'update', or 'delete' subcommands")
		os.Exit(1)
	}

	// Parse and execute subcommands
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		tags := strings.Split(*addTags, ",") // Split the tags into the slice
		addTodo(*addTitle, *addDesc, *addCategory, tags)
	case "list":
		listCmd.Parse(os.Args[2:])
		tags := strings.Split(*listTags, ",")
		listTodos(*listFilter, *listSort, *listCategory, tags)
	case "update":
		updateCmd.Parse(os.Args[2:])
		updateTodo(*updateID, *updateTitle, *updateDesc, *updateCompleted)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteAll {
			deleteAllTodos()
		} else if *deleteID == -1 {
			fmt.Println("Please introduce a todo ID higher or equal with 0.")
		} else {
			deleteTodo(*deleteID)
		}
	case "search":
		searchCmd.Parse(os.Args[2:])
		if *searchQuery == "" {
			fmt.Println("Please provide a query using -query flag.")
			os.Exit(1)
		}
		searchTodos(*searchQuery)
	case "stats":
		statsCmd.Parse(os.Args[2:])
		stats()
	default:
		fmt.Println("expected 'add', 'list', 'update', or 'delete' subcommands")
		os.Exit(1)
	}

}

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	Category    string
	Tags        []string
}

const filePath = "todo.csv"

func readTodos() ([]Todo, error) {
	file, err := os.Open(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil // If files doesn't exist, return an empty slice
		}
		return nil, err
	}
	defer file.Close()

	var todos []Todo
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records[1:] { // skipping the header row, so that's why we start from 1
		id, _ := strconv.Atoi(record[0])
		completed, _ := strconv.ParseBool(record[3])
		tags := strings.Split(record[5], ",") // Split tags by commas
		todos = append(todos, Todo{
			ID:          id,
			Title:       record[1],
			Description: record[2],
			Completed:   completed,
			Category:    record[4],
			Tags:        tags,
		})
	}

	return todos, nil
}

func writeTodos(todos []Todo) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	writer.Write([]string{"ID", "Title", "Description", "Completed", "Category", "Tags"})
	for _, todo := range todos {
		completed := strconv.FormatBool(todo.Completed)
		writer.Write([]string{
			strconv.Itoa(todo.ID),
			todo.Title,
			todo.Description,
			completed,
			todo.Category,
			strings.Join(todo.Tags, ","), // Join tags with commas
		})
	}

	return nil
}

func updateTodo(id int, title, desc string, completed bool) {
	todos, _ := readTodos()
	for i, todo := range todos {
		if todo.ID == id {
			// Update the fields only if new values are provided
			if title != "" {
				todos[i].Title = title
			}
			if desc != "" {
				todos[i].Description = desc
			}
			todos[i].Completed = completed
			writeTodos(todos) // Save updated todos back to the file
			fmt.Println("Todo updated successfully")
			return
		}
	}
	fmt.Println("Todo not found")
}

func addTodo(title, desc, category string, tags []string) {
	todos, _ := readTodos() // Read existing todos
	id := len(todos) + 1    // Assign a new ID
	todos = append(todos, Todo{
		ID:          id,
		Title:       title,
		Description: desc,
		Completed:   false,
		Category:    category,
		Tags:        tags, // Assign tags
	})
	writeTodos(todos) // Save the updated todos back to the file
	fmt.Println("Todo added successfully")
}

func deleteTodo(id int) {
	todos, _ := readTodos() // Read existing todos
	for i, todo := range todos {
		if todo.ID == id {
			// Remove the todo by slicking
			todos = append(todos[:i], todos[i+1:]...) // todos[:i] All elements before the index // todos[i+1:] All elements after the index
			writeTodos(todos)                         // Save the updated list back to the file
			fmt.Println("Todo deleted successfully!")
			return
		}
	}
	fmt.Println("Todo not found")
}

func deleteAllTodos() {
	writeTodos(nil)
}

func displayTodos(todos []Todo) {
	for _, todo := range todos {
		tags := strings.Join(todo.Tags, ", ") // Join the tags slice into a comma-separated string
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %v, Category: %s, Tags: [%s]\n",
			todo.ID, todo.Title, todo.Description, todo.Completed, todo.Category, tags)
	}
}

func listTodos(filter, sortBy, category string, tags []string) {
	todos, _ := readTodos() // Read existing todos

	// Apply filter
	if filter != "" {
		var filteredTodos []Todo
		for _, todo := range todos {
			if (filter == "completed" && todo.Completed) || (filter == "pending" && !todo.Completed) {
				filteredTodos = append(filteredTodos, todo)
			}
		}
		todos = filteredTodos
	}

	// Apply sorting
	if sortBy == "title" {
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].Title < todos[j].Title
		})
	} else if sortBy == "id" {
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].ID < todos[j].ID
		})
	}

	// Filter by category
	if category != "" && category != "all" {
		var categorizedTodos []Todo
		for _, todo := range todos {
			if strings.ToLower(todo.Category) == strings.ToLower(category) {
				categorizedTodos = append(categorizedTodos, todo)
			}
		}
		todos = categorizedTodos
	}

	// Filter by tags
	if len(tags) > 0 {
		var taggedTodos []Todo
		for _, todo := range todos {
			for _, tag := range tags {
				if contains(todo.Tags, tag) {
					taggedTodos = append(taggedTodos, todo)
					break
				}
			}
		}
		todos = taggedTodos
	}

	displayTodos(todos)
}

func searchTodos(query string) {
	todos, _ := readTodos() // Read existing todos

	// Filter todos by matching query in Title or Description
	var matchedTodos []Todo
	for _, todo := range todos {
		if strings.Contains(strings.ToLower(todo.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(todo.Description), strings.ToLower(query)) {
			matchedTodos = append(matchedTodos, todo)
		}
	}

	// Display matched todos
	if len(matchedTodos) == 0 {
		fmt.Println("No todos found matching the query.")
		return
	}

	displayTodos(matchedTodos)
}

func stats() {
	todos, _ := readTodos() // Read existing todos

	total := len(todos)
	completed := 0
	pending := 0
	categoryCounts := make(map[string]int) // Map to count todos by category

	for _, todo := range todos {
		if todo.Completed {
			completed++
		} else {
			pending++
		}
		categoryCounts[todo.Category]++ // Increment category count
	}

	// Calculate percentages
	completedPercentage := (float64(completed) / float64(total)) * 100
	pendingPercentage := (float64(pending) / float64(total)) * 100

	// Determine the most common category
	var mostCommonCategory string
	var maxCount int
	for category, count := range categoryCounts {
		if count > maxCount {
			mostCommonCategory = category
			maxCount = count
		}
	}

	// Display stats
	fmt.Println("Todo Statistics:")
	fmt.Printf("Total todos: %d\n", total)
	fmt.Printf("Completed: %d (%.2f%%)\n", completed, completedPercentage)
	fmt.Printf("Pending: %d (%.2f%%)\n", pending, pendingPercentage)

	fmt.Println("\nTodos by Category:")
	for categoy, count := range categoryCounts {
		fmt.Printf("%s: %d\n", categoy, count)
	}

	fmt.Printf("\nMost Common Category: %s (%d todos)\n", mostCommonCategory, maxCount)
}

func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if strings.ToLower(sliceItem) == strings.ToLower(item) {
			return true
		}
	}
	return false
}
