package handler

import (
	"hot-coffee/pkg/logger"
	"net/http"
)

func (myhandler *MyHandler) errorHandling(w http.ResponseWriter, message string, status int) {
	logger.MyLogger.Debug("Activation of Function", "Layer", "Handler", "Function", "errorHandling")
	w.WriteHeader(status)
	err := myhandler.Representation.ErrorRepresentation(w, status, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	logger.MyLogger.Debug("End of Function", "Layer", "Handler", "Function", "errorHandling")

}
