package handler

import "SimpleCoffee/internal/domain/module"

type MyHandler struct {
	Service module.Service
}

func NewMyHandler(service module.Service) *MyHandler {
	return &MyHandler{
		Service: service,
	}
}
