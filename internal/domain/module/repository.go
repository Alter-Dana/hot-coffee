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
	UpdataInventory(inventory *entity.InventoryItem) error
}
type menuRepo interface {
}
