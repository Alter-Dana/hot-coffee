package handler

import (
	"net/http"
)

func (myhandler *MyHandler) errorHandling(w http.ResponseWriter, message string, status int) {
	err := myhandler.Representation.ErrorRepresentation(w, status, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(status)

}
