package handler

import (
	"SimpleCoffee/pkg/logger"
	"net/http"
)

func (myhandler *MyHandler) errorHandling(w http.ResponseWriter, message string, status int) {
	logger.MyLogger.Debug("Activation of Function", "Layer", "Handler", "Function", "errorHandling")
	err := myhandler.Representation.ErrorRepresentation(w, status, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(status)
	logger.MyLogger.Debug("End of Function", "Layer", "Handler", "Function", "errorHandling")

}
