package dal

import (
	"SimpleCoffee/internal/domain"
	"SimpleCoffee/internal/domain/entity"
	"SimpleCoffee/pkg/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var mapInventory map[string]*entity.InventoryItem = make(map[string]*entity.InventoryItem)

func (rp *Repository) CreateInventory(inventory *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "CreateInventory")

	mapInventory[inventory.IngredientID] = inventory

	dir, err := openDirectory(rp.Directory)
	if err != nil {
		return err
	}
	defer dir.Close()
	// Try to append the new object to the JSON file
	if err := appendToJSONInventoryFile(rp.Directory, rp.FileInventory, inventory); err != nil {
		return err
	}

	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "CreateInventory")
	return nil
}

func (rp *Repository) GetCertainInventory(inventoryID string) *entity.InventoryItem {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "GetCertainInventory")

	inventory, ok := mapInventory[inventoryID]

	if !ok {
		return nil
	}
	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "GetCertainInventory")
	return inventory
}

func (rp *Repository) GetAllInventory() ([]*entity.InventoryItem, error) {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "GetAllInventory")
	// read inventory json file  and unmarshall json into inventory struct, handle errors and log them
	// return
	//
}

func openDirectory(path string) (*os.File, error) {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "openDirectory")

	dir, err := os.Open(path)

	if err != nil {
		if os.IsNotExist(err) {
			logger.MyLogger.Error("Failed to open the directory, because it does not exist", "Layer", "Repository", "Function", "openDirectory", "error", err.Error())
			return nil, domain.ErrInternalServer
		}
		logger.MyLogger.Error("Failed to open the directory, because it does not exist", "Layer", "Repository", "Function", "openDirectory", "error", err.Error())
		return nil, domain.ErrInternalServer
	}

	// Ensure it's actually a directory
	info, err := dir.Stat()
	if err != nil {
		logger.MyLogger.Error("Failed to identify getting information in order to know whether it is actually a directory or not", "Layer", "Repository", "Function", "openDirectory", "error", err.Error())
		err = dir.Close()
		if err != nil {
			logger.MyLogger.Error("Failed to close the directory, after opening it", "Layer", "Repository", "Function", "openDirectory", "error", err.Error())
		}
		return nil, domain.ErrInternalServer
	}
	if !info.IsDir() {

		logger.MyLogger.Error(fmt.Sprintf("Such kind of name does exist:%s, but it is not Directory", path), "Layer", "Repository", "Function", "openDirectory", "error", err.Error())

		err = dir.Close()
		if err != nil {
			logger.MyLogger.Error("Failed to close the directory, after opening it", "Layer", "Repository", "Function", "openDirectory", "error", err.Error())
		}

		return nil, domain.ErrInternalServer
	}

	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "openDirectory")

	return dir, nil
}

func appendToJSONInventoryFile(directory, fileName string, newInventory *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "appendToJSONInventoryFile")

	filePath := filepath.Join(directory, fileName)

	// Check if the JSON file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			logger.MyLogger.Error(fmt.Sprintf("Failed to open JSON file '%s' does not exist", filePath), "Layer", "Repository", "Function", "appendToJSONInventoryFile", "error", err.Error())
			return domain.ErrInternalServer
		}
		logger.MyLogger.Error("Failed to check JSON file", "Layer", "Repository", "Function", "appendToJSONInventoryFile", "error", err.Error())
		return domain.ErrInternalServer
	}

	// Read the existing JSON file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.MyLogger.Error("Failed to read JSON file", "Layer", "Repository", "Function", "appendToJSONInventoryFile", "error", err.Error())
		return domain.ErrInternalServer
	}

	// Initialize a slice to hold Person objects
	var inventoryItems []entity.InventoryItem

	// Unmarshal only if the file is not empty
	if len(fileData) > 0 {
		if err := json.Unmarshal(fileData, &inventoryItems); err != nil {
			logger.MyLogger.Error("Failed to unmarshall JSON", "Layer", "Repository", "Function", "appendToJSONInventoryFile", "error", err.Error())
			return domain.ErrInternalServer
		}
	}

	// Append new data
	inventoryItems = append(inventoryItems, *newInventory)

	// Marshal updated JSON
	updatedJSON, err := json.MarshalIndent(inventoryItems, "", "  ")
	if err != nil {
		logger.MyLogger.Error("Failed to marshall JSON", "Layer", "Repository", "Function", "appendToJSONInventoryFile", "error", err.Error())
		return domain.ErrInternalServer
	}

	// Write back the updated JSON
	if err := ioutil.WriteFile(filePath, updatedJSON, 0644); err != nil {
		logger.MyLogger.Error("Failed to write JSON file", "Layer", "Repository", "Function", "appendToJSONInventoryFile", "error", err.Error())
		return domain.ErrInternalServer
	}

	logger.MyLogger.Info("\tfmt.Sprintf(\"Successfully appended new data to '%s'\\n\", filePath)\n")
	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "appendToJSONInventoryFile")
	return nil
}
