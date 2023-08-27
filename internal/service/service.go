package service

type Service struct {
	US UserService
	OS OrderService
	IS ItemService
}

func New(us UserService, os OrderService, is ItemService) *Service {
	return &Service{
		US: us,
		OS: os,
		IS: is,
	}
}
