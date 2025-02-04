package cmd

import (
	"SimpleCoffee/configs"
	"SimpleCoffee/internal/dal"
	"SimpleCoffee/internal/handler"
	"SimpleCoffee/internal/representation/myjson"
	"SimpleCoffee/internal/service"
	"SimpleCoffee/pkg/logger"
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
	logger.MyLogger.Info("Starting server at the port: ", *conf.Port)
	fmt.Println("The server is running at: http://localhost:", *conf.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *conf.Port), router))

}
