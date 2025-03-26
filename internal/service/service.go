package service

import "hot-coffee/internal/domain/module"

type MyService struct {
	MyDB module.Repository
}

func NewMyService(myDB module.Repository) *MyService {
	return &MyService{
		MyDB: myDB,
	}
}
