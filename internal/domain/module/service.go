package module

import "SimpleCoffee/internal/domain/entity"

type Service interface {
	inventory
}

type inventory interface {
	CreateInventory(inventoryItem *entity.InventoryItem) error
}
