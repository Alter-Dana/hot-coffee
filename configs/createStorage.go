package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func (c *Config) CreateStorage() error {
	absPath, err := filepath.Abs(*c.Dir)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path: %v", err)
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(absPath, 0o777); err != nil {
			return fmt.Errorf("error in storage creation: failed to create root directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error accessing the directory: %w", err)
	} else if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", absPath)
	}

	// To ensure that this folder is "writable"
	testFile := filepath.Join(absPath, ".testfile")
	file, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("problem with writing in the directory: %w", err)
	}

	file.Close()
	os.Remove(testFile)

	err = c.CreateJSONs()
	if err != nil {
		return errors.New("error creating json storage files")
	}
	return nil
}

// Work on this later
func (c *Config) CreateJSONs() error {
	files := []string{"inventory_repo.json", "menu_repo.json", "order_repo.json"}

	for _, file := range files {
		filePath := filepath.Join(*c.Dir, file)

		// Check if the file already exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Create the file
			file, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create %s: %v", filePath, err)
			}
			defer file.Close()

			// Write an empty JSON array `[]`
			encoder := json.NewEncoder(file)
			if err := encoder.Encode([]interface{}{}); err != nil {
				return fmt.Errorf("failed to initialize %s: %v", filePath, err)
			}
		}
	}

	return nil
}
