package service

import (
	"hot-coffee/internal/domain"
	"hot-coffee/internal/domain/entity"
	"hot-coffee/pkg/logger"
)

func (app *MyService) CreateInventory(inventoryItem *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Service", "Function", "CreateInventory")

	if inventoryItem.Name == "" || inventoryItem.Quantity <= 0 || inventoryItem.Unit == "" || inventoryItem.IngredientID == "" {
		logger.MyLogger.Info("The value of the sent body is invalid")
		return domain.ErrInvalidInventory
	}

	dbInventory := app.MyDB.GetCertainInventory(inventoryItem.IngredientID)
	if dbInventory != nil {
		if inventoryItem.Name == dbInventory.Name && inventoryItem.Unit == dbInventory.Unit {
			inventoryItem.Quantity += dbInventory.Quantity
			err := app.MyDB.UpdateInventory(inventoryItem)
			if err != nil {
				return err
			}
			return nil
		} else {
			logger.MyLogger.Info("The value of the sent body is invalid")
			return domain.ErrInvalidInventory
		}
	}

	err := app.MyDB.CreateInventory(inventoryItem)
	if err != nil {
		return err
	}

	logger.MyLogger.Debug("End of function", "Layer", "Service", "Function", "CreateInventory")
	return nil

}

func (app *MyService) GetInventory() ([]entity.InventoryItem, error) {
	logger.MyLogger.Debug("Activation of function", "Layer", "Service", "Function", "GetInventory")
	return nil, nil
}

func (app *MyService) UpdateInventory(inventoryItem *entity.InventoryItem) error {
	logger.MyLogger.Debug("Activation of function", "Layer", "Service", "Function", "UpdateInventory")
	if inventoryItem.IngredientID == "" {
		logger.MyLogger.Info("The value of the sent body is invalid")
		return domain.ErrInvalidInventory
	}

	dbInventory := app.MyDB.GetCertainInventory(inventoryItem.IngredientID)
	if dbInventory == nil {
		logger.MyLogger.Info("The value of the sent body is invalid")
		return domain.ErrInvalidInventory
	}

	// Do forget to update the related map data

	if inventoryItem.Name != "" {
		dbInventory.Name = inventoryItem.Name
	}
	if inventoryItem.Quantity > 0 {
		dbInventory.Quantity = inventoryItem.Quantity
	}
	if inventoryItem.Unit != "" {
		dbInventory.Unit = inventoryItem.Unit
	}

	err := app.MyDB.UpdateInventory(dbInventory)
	if err != nil {
		return err
	}

	return nil
}
