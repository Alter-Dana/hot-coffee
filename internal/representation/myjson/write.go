package myjson

import (
	"SimpleCoffee/internal/domain/entity"
	"SimpleCoffee/pkg/logger"
	"encoding/json"
	"net/http"
)

func (representation *JSONRepresentation) ConvertInventoryToResponse(w http.ResponseWriter, inventories []entity.InventoryItem) error {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&inventories); err != nil {
		logger.MyLogger.Error("Failed to encode the structure into JSON", "Layer", "Representation", "Function", "ConvertInventoryToResponse", "error", err)
		return err
	}
	return nil
}
func (representation *JSONRepresentation) ErrorRepresentation(w http.ResponseWriter, status int, message string) error {
	errorData := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Error:   http.StatusText(status),
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&errorData); err != nil {
		logger.MyLogger.Error("Failed to encode the structure into JSON", "Layer", "Representation", "Function", "ErrorRepresentation", "error", err)
		return err
	}

	return nil
}
