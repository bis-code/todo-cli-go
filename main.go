package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Add command flags
	addTitle := addCmd.String("title", "", "The title of the todo")
	addDesc := addCmd.String("desc", "", "Description of the todo")

	// Update command flags
	updateID := updateCmd.Int("id", 0, "ID of the todo to update")
	updateTitle := updateCmd.String("title", "", "Updated title")
	updateDesc := updateCmd.String("desc", "", "Updated description")
	updateCompleted := updateCmd.Bool("completed", false, "Update completed task")

	// Delete command flag
	deleteID := deleteCmd.Int("id", 0, "ID of the todo to delete")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', 'update', or 'delete' subcommands")
		os.Exit(1)
	}

	// Parse and execute subcommands
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		addTodo(*addTitle, *addDesc)
	case "list":
		listCmd.Parse(os.Args[2:])
		listTodos()
	case "update":
		updateCmd.Parse(os.Args[2:])
		updateTodo(*updateID, *updateTitle, *updateDesc, *updateCompleted)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		deleteTodo(*deleteID)
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
		todos = append(todos, Todo{ID: id, Title: record[1], Description: record[2], Completed: completed})
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
	writer.Write([]string{"ID", "Title", "Description", "Completed"})
	for _, todo := range todos {
		completed := strconv.FormatBool(todo.Completed)
		writer.Write([]string{strconv.Itoa(todo.ID), todo.Title, todo.Description, completed})
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

func addTodo(title, desc string) {
	todos, _ := readTodos() // Read existing todos
	id := len(todos) + 1    // Assign a new ID
	todos = append(todos, Todo{ID: id, Title: title, Description: desc, Completed: false})
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

func listTodos() {
	todos, _ := readTodos() // Read existing todos
	for _, todo := range todos {
		fmt.Printf("ID: %d, Title %s, Description: %s, Completed: %v\n", todo.ID, todo.Title, todo.Description, todo.Completed)
	}
}
