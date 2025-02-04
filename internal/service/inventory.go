package service

import (
	"SimpleCoffee/internal/domain"
	"SimpleCoffee/internal/domain/entity"
	"SimpleCoffee/pkg/logger"
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
			err := app.MyDB.UpdateInventory(inventoryItem.IngredientID, inventoryItem.Quantity, "addition")
			if err != nil {

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

}
