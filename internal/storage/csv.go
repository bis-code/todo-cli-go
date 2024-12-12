package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"todo-cli/internal/models"
)

const CSVFilePath = "todos.csv"

func ReadTodosFromCSV() ([]models.Todo, error) {
	file, err := os.Open(CSVFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var todos []models.Todo
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records[1:] { // Skip header row
		id, _ := strconv.Atoi(record[0])
		completed, _ := strconv.ParseBool(record[3])
		tags := strings.Split(record[5], ",")
		todos = append(todos, models.Todo{
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

func WriteTodosToCSV(todos []models.Todo) error {
	file, err := os.Create(CSVFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"ID", "Title", "Description", "Completed", "Category", "Tags"})
	for _, todo := range todos {
		completed := strconv.FormatBool(todo.Completed)
		tags := strings.Join(todo.Tags, ",")
		writer.Write([]string{
			strconv.Itoa(todo.ID),
			todo.Title,
			todo.Description,
			completed,
			todo.Category,
			tags,
		})
	}

	return nil
}
