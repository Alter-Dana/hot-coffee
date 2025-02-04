package module

import "SimpleCoffee/internal/domain/entity"

type Repository interface {
	orderRepo
	menuRepo
	inventoryRepo
}

type orderRepo interface {
}

type inventoryRepo interface {
	GetCertainInventory(inventoryID string) *entity.InventoryItem
	CreateInventory(inventory *entity.InventoryItem) error
	UpdateInventory(inventory *entity.InventoryItem) error
	GetAllInventory() ([]entity.InventoryItem, error)
}
type menuRepo interface {
}
