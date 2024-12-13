package main

import (
	"flag"
	"fmt"
	"os"

	"todo-cli/cmd" // Import your command package
	"todo-cli/internal/storage"
)

func main() {
	// Initialize database and apply migrations
	storage.InitDB()
	defer storage.CloseDB()

	// Ensure a subcommand is provided
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse subcommands
	switch os.Args[1] {
	case "migrate":
		runMigrations()
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "The title of the todo")
		desc := addCmd.String("desc", "", "Description of the todo")
		category := addCmd.String("category", "General", "Category of the todo")
		tags := addCmd.String("tags", "", "Comma-separated tags for the todo")
		addCmd.Parse(os.Args[2:])
		if *title == "" {
			fmt.Println("Error: The 'title' flag is required for the 'add' subcommand.")
			addCmd.Usage()
			os.Exit(1)
		}
		cmd.Add(*title, *desc, *category, *tags)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filter := listCmd.String("filter", "", "Filter todos by 'completed' or 'pending'")
		sortBy := listCmd.String("sort", "id", "Sort todos by 'id' or 'title'")
		category := listCmd.String("category", "all", "Filter todos by category")
		tags := listCmd.String("tags", "", "Comma-separated tags to filter by")
		listCmd.Parse(os.Args[2:])
		cmd.List(*filter, *sortBy, *category, *tags)

	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.Int("id", 0, "ID of the todo to update")
		title := updateCmd.String("title", "", "Updated title")
		desc := updateCmd.String("desc", "", "Updated description")
		completed := updateCmd.Bool("completed", false, "Mark the todo as completed")
		updateCmd.Parse(os.Args[2:])
		if *id <= 0 {
			fmt.Println("Error: A valid 'id' is required for the 'update' subcommand.")
			updateCmd.Usage()
			os.Exit(1)
		}
		cmd.Update(*id, *title, *desc, *completed)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", -1, "ID of the todo to delete")
		deleteAll := deleteCmd.Bool("all", false, "Delete all todos")
		deleteCmd.Parse(os.Args[2:])
		if *deleteAll {
			cmd.Delete(0, true)
		} else if *id >= 0 {
			cmd.Delete(*id, false)
		} else {
			fmt.Println("Error: Provide a valid 'id' or use the '-all' flag for the 'delete' subcommand.")
			deleteCmd.Usage()
			os.Exit(1)
		}

	case "search":
		searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
		query := searchCmd.String("query", "", "Search query for todos")
		searchCmd.Parse(os.Args[2:])
		if *query == "" {
			fmt.Println("Error: The 'query' flag is required for the 'search' subcommand.")
			searchCmd.Usage()
			os.Exit(1)
		}
		cmd.Search(*query)

	case "stats":
		cmd.Stats()

	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

// printUsage prints a help message with available subcommands
func printUsage() {
	fmt.Println("Usage: todo-cli <subcommand> [options]")
	fmt.Println("\nAvailable subcommands:")
	fmt.Println("  add     Add a new todo")
	fmt.Println("  list    List todos")
	fmt.Println("  update  Update an existing todo")
	fmt.Println("  delete  Delete a todo by ID or all todos")
	fmt.Println("  search  Search todos by a query")
	fmt.Println("  stats   Display statistics about todos")
	fmt.Println("\nRun 'todo-cli <subcommand> -h' for more information about a specific command.")
}

func runMigrations() {
	err := storage.RunMigrations()
	if err != nil {
		fmt.Println("Migrations failed:", err)
		os.Exit(1)
	}
}
