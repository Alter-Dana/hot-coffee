package module

import "hot-coffee/internal/domain/entity"

type Service interface {
	inventory
}

type inventory interface {
	CreateInventory(inventoryItem *entity.InventoryItem) error
	GetInventory() ([]entity.InventoryItem, error)
}
