package handler

import "net/http"

func (myHandler *MyHandler) home(w http.ResponseWriter, r *http.Request) {

	myHandler.errorHandling(w, "Not Found", http.StatusNotFound)

}
