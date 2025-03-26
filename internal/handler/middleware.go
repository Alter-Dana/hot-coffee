package handler

import (
	"hot-coffee/pkg/logger"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		logger.MyLogger.Info(fmt.Sprintf("Method:[%v], URL_Path: %v, Remote_Address: %v\n", r.Method, r.URL.Path, r.RemoteAddr))
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		// Here must be reponse's status code
		logger.MyLogger.Info(fmt.Sprintf("The End of the client request, and its Duration:%v\n", duration))
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.MyLogger.Error(fmt.Sprintf("Panic occurred:%v\n%v\n", err, string(debug.Stack())))
				//serverError(w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
