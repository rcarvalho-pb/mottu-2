package controller

import (
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/helper"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

type AuthController struct {
	service *service.Service
}

func newAuthController(serv *service.Service) AuthController {
	return AuthController{service: serv}
}

func (ac *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var authRequest model.AuthRequest
	if err := helper.ReadJson(w, r, &authRequest); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	tokenString, err := ac.service.AuthService.GenerateToken(&authRequest)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, tokenString)
}
