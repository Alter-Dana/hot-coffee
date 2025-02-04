package dal

import (
	"SimpleCoffee/internal/domain/entity"
	"SimpleCoffee/pkg/logger"
)

var mapInventory map[string]*entity.InventoryItem = make(map[string]*entity.InventoryItem)

func (rp *Repository) CreateInventory(inventory *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Repository", "Function", "CreateInventory")
	mapInventory[inventory.IngredientID] = inventory

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
