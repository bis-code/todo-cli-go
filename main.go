package main

import (
	"flag"
	"fmt"
	"os"
	"todo-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', 'update', 'delete', 'search', or 'stats' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "The title of the todo")
		desc := addCmd.String("desc", "", "Description of the todo")
		category := addCmd.String("category", "General", "Category of the todo")
		tags := addCmd.String("tags", "", "Comma-separated tags for the todo")
		addCmd.Parse(os.Args[2:])
		if *title == "" {
			fmt.Println("Title is required.")
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
		completed := updateCmd.Bool("completed", false, "Mark as completed")
		updateCmd.Parse(os.Args[2:])
		if *id <= 0 {
			fmt.Println("Please provide a valid ID using -id flag.")
			os.Exit(1)
		}
		cmd.Update(*id, *title, *desc, *completed)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", -1, "ID of the todo to delete")
		deleteAll := deleteCmd.Bool("all", false, "Deletes all todos")
		deleteCmd.Parse(os.Args[2:])
		if *deleteAll {
			cmd.Delete(0, true)
		} else if *id > 0 {
			cmd.Delete(*id, false)
		} else {
			fmt.Println("Please provide a valid ID using -id flag or use -all to delete all todos.")
			os.Exit(1)
		}

	case "search":
		searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
		query := searchCmd.String("query", "", "Search query for todos")
		searchCmd.Parse(os.Args[2:])
		if *query == "" {
			fmt.Println("Please provide a query using -query flag.")
			os.Exit(1)
		}
		cmd.Search(*query)

	case "stats":
		statsCmd := flag.NewFlagSet("stats", flag.ExitOnError)
		statsCmd.Parse(os.Args[2:])
		cmd.Stats()

	default:
		fmt.Println("Expected 'add', 'list', 'update', 'delete', 'search', or 'stats' subcommands")
		os.Exit(1)
	}
}
