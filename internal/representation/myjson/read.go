package myjson

import (
	"hot-coffee/internal/domain"
	"hot-coffee/internal/domain/entity"
	"hot-coffee/pkg/logger"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (representation *JSONRepresentation) ConvertToInventoryObject(r *http.Request) (*entity.InventoryItem, error) {
	var obj *entity.InventoryItem = &entity.InventoryItem{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&obj); err != nil {

		if errors.Is(err, io.EOF) || errors.Is(err, &json.SyntaxError{}) || errors.Is(err, &json.UnmarshalTypeError{}) || err.Error() == "unknown field" {
			logger.MyLogger.Info("The incoming request's body is invalid", "Layer", "Representation", "Function", "ConvertToInventoryObject", "error", err.Error())
			return nil, err
		} else {
			logger.MyLogger.Error("Failed to decode the incoming request's body", "Layer", "Representation", "Function", "ConvertToInventoryObject", "error", err.Error())
			return nil, domain.ErrInternalServer
		}
	}

	return obj, nil
}

func (representation *JSONRepresentation) ConvertToMenuObject(r *http.Request) (*entity.MenuItem, error) {
	var obj *entity.MenuItem = &entity.MenuItem{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&obj); err != nil {

		if errors.Is(err, io.EOF) || errors.Is(err, &json.SyntaxError{}) || errors.Is(err, &json.UnmarshalTypeError{}) || err.Error() == "unknown field" {
			logger.MyLogger.Info("The incoming request's body is invalid", "Layer", "Representation", "Function", "ConvertToMenuObject", "error", err.Error())
			return nil, err
		} else {
			logger.MyLogger.Error("Failed to decode the incoming request's body", "Layer", "Representation", "Function", "ConvertToMenuObject", "error", err.Error())
			return nil, domain.ErrInternalServer
		}
	}

	return obj, nil
}

func (representation *JSONRepresentation) ConvertToOrderObject(r *http.Request) (*entity.OrderItem, error) {
	var obj *entity.OrderItem = &entity.OrderItem{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&obj); err != nil {

		if errors.Is(err, io.EOF) || errors.Is(err, &json.SyntaxError{}) || errors.Is(err, &json.UnmarshalTypeError{}) || err.Error() == "unknown field" {
			logger.MyLogger.Info("The incoming request's body is invalid", "Layer", "Representation", "Function", "ConvertToOrderObject", "error", err.Error())
			return nil, err
		} else {
			logger.MyLogger.Error("Failed to decode the incoming request's body", "Layer", "Representation", "Function", "ConvertToOrderObject", "error", err.Error())
			return nil, domain.ErrInternalServer
		}
	}

	return obj, nil
}
