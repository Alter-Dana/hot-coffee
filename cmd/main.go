package main

import (
	"hot-coffee/configs"
	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/representation/myjson"
	"hot-coffee/internal/service"
	"hot-coffee/pkg/logger"
	"fmt"
	"log"
	"net/http"
)

func main() {

	logger.MyLogger = logger.GetLoggerObject("../logging/logging.log")

	conf, err := configs.NewConfiguration()
	if err != nil {
		logger.MyLogger.Error("Failed to configure", "error", err)
		return
	}

	myRepo := dal.NewRepository(*conf.Dir)
	myService := service.NewMyService(myRepo)
	myRepresentation := myjson.NewRepresentation()
	myHandler := handler.NewMyHandler(myService, myRepresentation)
	router := myHandler.InitRouter()
	logger.MyLogger.Info(fmt.Sprintf("Starting server at the port: %v", *conf.Port))
	fmt.Println(fmt.Sprintf("The server is running at: http://localhost:%v", *conf.Port))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *conf.Port), router))

}
