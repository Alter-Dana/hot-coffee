package module

import (
	"SimpleCoffee/internal/domain/entity"
	"net/http"
)

type Representation interface {
	inventoryRepresentation
	errorRepresentation
}

type orderRepresentation interface {
}
type inventoryRepresentation interface {
	ConvertToInventoryObject(r *http.Request) (*entity.InventoryItem, error)
	ConvertInventoryToResponse(w http.ResponseWriter, inventories []entity.InventoryItem) error
}

type menuRepresentation interface{}

type errorRepresentation interface {
	ErrorRepresentation(w http.ResponseWriter, status int, message string) error
}
