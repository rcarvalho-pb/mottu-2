package controller

import "github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"

type Controller struct {
	UserController UserController
	AuthController AuthController
}

func New() Controller {
	service := service.New()
	return Controller{
		UserController: newUserController(service),
		AuthController: newAuthController(service),
	}
}
