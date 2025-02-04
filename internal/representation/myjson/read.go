package myjson

import (
	"SimpleCoffee/internal/domain/entity"
	"SimpleCoffee/pkg/logger"
	"encoding/json"
	"net/http"
)

func (representation *JSONRepresentation) ConvertToInventoryObject(r *http.Request) (*entity.InventoryItem, error) {
	var obj *entity.InventoryItem = &entity.InventoryItem{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&obj); err != nil {
		logger.MyLogger.Error("Failed to decode the incoming request's body", "Layer", "Representation", "Function", "ConvertToInventoryObject", "error", err.Error())
		return nil, err
	}

	return obj, nil
}

func (representation *JSONRepresentation) ConvertToMenuObject(r *http.Request) (*entity.MenuItem, error) {
	var obj *entity.MenuItem = &entity.MenuItem{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&obj); err != nil {
		logger.MyLogger.Error("Failed to decode the incoming request's body", "Layer", "Representation", "Function", "ConvertToInventoryObject", "error", err.Error())
		return nil, err
	}

	return obj, nil
}

func (representation *JSONRepresentation) ConvertToOrderObject(r *http.Request) (*entity.OrderItem, error) {
	var obj *entity.OrderItem = &entity.OrderItem{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&obj); err != nil {
		logger.MyLogger.Error("Failed to decode the incoming request's body", "Layer", "Representation", "Function", "ConvertToInventoryObject", "error", err.Error())
		return nil, err
	}

	return obj, nil
}
