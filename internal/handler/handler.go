package handler

import "SimpleCoffee/internal/domain/module"

type MyHandler struct {
	Service        module.Service
	Representation module.Representation
}

func NewMyHandler(service module.Service, representation module.Representation) *MyHandler {
	return &MyHandler{
		Service:        service,
		Representation: representation,
	}
}
