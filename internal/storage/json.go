package storage

import (
	"encoding/json"
	"os"

	"todo-cli/internal/models"
)

const JSONFilePath = "todos.json"

func ReadTodosFromJSON() ([]models.Todo, error) {
	file, err := os.Open(JSONFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var todos []models.Todo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func WriteTodosToJSON(todos []models.Todo) error {
	file, err := os.Create(JSONFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON for readability
	return encoder.Encode(todos)
}
