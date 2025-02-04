package myjson

import (
	"SimpleCoffee/pkg/logger"
	"encoding/json"
	"net/http"
)

func (representation *JSONRepresentation) ErrorRepresentation(w http.ResponseWriter, status int, message string) error {
	errorData := struct {
		error   string
		message string
	}{
		error:   http.StatusText(status),
		message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(errorData); err != nil {
		logger.MyLogger.Error("Failed to encode the structure into JSON", "Layer", "Representation", "Function", "ErrorRepresentation", "error", err)
		return err
	}
	return nil
}
