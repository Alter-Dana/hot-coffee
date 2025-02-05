package handler

import (
	"SimpleCoffee/pkg/logger"
	"net/http"
)

func (myHandler *MyHandler) home(w http.ResponseWriter, r *http.Request) {
	logger.MyLogger.Debug("Activation of Function", "Layer", "Handler", "Function", "home")
	myHandler.errorHandling(w, "Not Found", http.StatusNotFound)
	logger.MyLogger.Debug("End of Function", "Layer", "Handler", "Function", "home")

}
