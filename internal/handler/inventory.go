package handler

import (
	"SimpleCoffee/internal/domain"
	"SimpleCoffee/pkg/logger"
	"errors"
	"net/http"
)

func (myhandler *MyHandler) inventory(w http.ResponseWriter, r *http.Request) {
	logger.MyLogger.Debug("Activation of function", "Layer", "Handler", "Function", "inventory")

	if r.URL.Path != "/inventory" {
		logger.MyLogger.Info("incoming request's path is invalid", "endpoint", "/inventory", "requestPath", r.URL.Path)
		myhandler.errorHandling(w, "There is no such endpoint:"+r.URL.Path, http.StatusNotFound)
		return
	}
	contentType := r.Header.Get("Content-Type")

	switch r.Method {
	case "GET":
		// kamila
		/*

			check for body
			get array of inventory items from json files
			inventory item struct is wrapped in representation here
			call convert

		*/

	case "POST":
		if contentType != "application/myjson" {
			logger.MyLogger.Info("incoming request's Content-type is invalid", "endpoint", "/inventory", "requestContentType", contentType)
			myhandler.errorHandling(w, "Content-Type should be application/myjson", http.StatusBadRequest)
			return
		}

		requestedInventory, err := myhandler.Representation.ConvertToInventoryObject(r)
		if errors.Is(err, domain.ErrInternalServer) {
			myhandler.errorHandling(w, err.Error(), http.StatusInternalServerError)
			return
		} else if err != nil {
			myhandler.errorHandling(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = myhandler.Service.CreateInventory(requestedInventory)
		if errors.Is(err, domain.ErrInternalServer) {
			myhandler.errorHandling(w, err.Error(), http.StatusInternalServerError)
			return
		} else if errors.Is(err, domain.ErrExistID) || errors.Is(err, domain.ErrInvalidInventory) {
			myhandler.errorHandling(w, err.Error(), http.StatusBadRequest)
			return
		}

	default:
		logger.MyLogger.Info("incoming request's method is invalid", "endpoint", "/inventory", "requestMethod", r.Method)
		myhandler.errorHandling(w, "incoming request's method is invalid", http.StatusMethodNotAllowed)
		return
	}

}

func (myhandler *MyHandler) specificInventory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/inventory/" {
		//myhandler.errorHandling()
		return
	}

	if r.Method != "GET" && r.Method != "PUT" && r.Method != "DELETE" {
		return
	}

	contentType := r.Header.Get("Content-Type")

	if r.Method == "GET" {

	} else if r.Method == "PUT" {
		if contentType != "application/myjson" {
			logger.MyLogger.Info("incoming request's Content-type is invalid", "endpoint", "/inventory/", "requestContentType", contentType)
			myhandler.errorHandling(w, "Content-Type should be application/myjson", http.StatusBadRequest)
			return
		}
	} else if r.Method == "DELETE" {

	}

}
