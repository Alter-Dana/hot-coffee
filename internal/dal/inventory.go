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

func (rp *Repository) GetAllInventory() ([]entity.InventoryItem, error) {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "GetAllInventory")

	var inventories []entity.InventoryItem

	dir, err := openDirectory(rp.Directory)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	inventories, err = retrieveInventoriesFromJson(rp.Directory + rp.FileInventory)
	if err != nil {
		return nil, err
	}
	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "GetAllInventory")
	return inventories, nil
}

func (rp *Repository) UpdateInventory(newInventory *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "UpdateInventory")
	dbInventories, err := rp.GetAllInventory()
	if err != nil {
		return err
	}
	for i := 0; i < len(dbInventories); i++ {
		if newInventory.IngredientID == dbInventories[i].IngredientID {
			dbInventories[i].Name = newInventory.Name
			dbInventories[i].Quantity = newInventory.Quantity
			dbInventories[i].Unit = newInventory.Unit
			mapInventory[newInventory.IngredientID] = &dbInventories[i]
			break
		}
	}

	err = writeInventoryJsonDataToFile(rp.Directory+rp.FileInventory, dbInventories)
	if err != nil {
		return err
	}

	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "UpdateInventory")
	return nil
}

// -----------------------------------------------
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

func retrieveRawDataFromJson(filePath string) ([]byte, error) {

	// Check if the JSON file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			logger.MyLogger.Error(fmt.Sprintf("Failed to open JSON file '%s' does not exist", filePath), "Layer", "Repository", "Function", "retrieveRawDataFromJson", "error", err.Error())
			return nil, domain.ErrInternalServer
		}
		logger.MyLogger.Error("Failed to check JSON file", "Layer", "Repository", "Function", "retrieveRawDataFromJson", "error", err.Error())
		return nil, domain.ErrInternalServer
	}

	// Read the existing JSON file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.MyLogger.Error("Failed to read JSON file", "Layer", "Repository", "Function", "retrieveRawDataFromJson", "error", err.Error())
		return nil, domain.ErrInternalServer
	}

	return fileData, nil
}

func retrieveInventoriesFromJson(filePath string) ([]entity.InventoryItem, error) {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "retrieveInventoriesFromJson")

	inventories := []entity.InventoryItem{}

	fileData, err := retrieveRawDataFromJson(filePath)
	if err != nil {
		return nil, err
	}

	if len(fileData) > 0 {
		if err := json.Unmarshal(fileData, &inventories); err != nil {
			logger.MyLogger.Error("Failed to unmarshall JSON", "Layer", "Repository", "Function", "retrieveInventoriesFromJson", "error", err.Error())
			return nil, domain.ErrInternalServer
		}
	}

	logger.MyLogger.Debug("End of function", "Layer", "Repository", "retrieveInventoriesFromJson")

	return inventories, nil
}

func appendToJSONInventoryFile(directory, fileName string, newInventory *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "appendToJSONInventoryFile")

	filePath := filepath.Join(directory, fileName)

	// Initialize a slice to hold Person objects
	inventories, err := retrieveInventoriesFromJson(filePath)
	if err != nil {
		return err
	}

	// Append new data
	inventories = append(inventories, *newInventory)

	// Marshal updated JSON
	err = writeInventoryJsonDataToFile(filePath, inventories)
	if err != nil {
		return err
	}

	logger.MyLogger.Info("\tfmt.Sprintf(\"Successfully appended new data to '%s'\\n\", filePath)\n")
	logger.MyLogger.Debug("End of function", "Layer", "Repository", "Function", "appendToJSONInventoryFile")
	return nil
}

func writeInventoryJsonDataToFile(filePath string, inventories []entity.InventoryItem) error {

	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "writeInventoryJsonDataToFile")

	updatedJSON, err := json.MarshalIndent(inventories, "", "  ")
	if err != nil {
		logger.MyLogger.Error("Failed to marshall JSON", "Layer", "Repository", "Function", "writeInventoryJsonDataToFile", "error", err.Error())
		return domain.ErrInternalServer
	}

	// Write back the updated JSON
	if err := ioutil.WriteFile(filePath, updatedJSON, 0644); err != nil {
		logger.MyLogger.Error("Failed to write JSON file", "Layer", "Repository", "Function", "writeInventoryJsonDataToFile", "error", err.Error())
		return domain.ErrInternalServer
	}

	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "writeInventoryJsonDataToFile")
	return nil
}
