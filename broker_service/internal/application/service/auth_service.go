package service

import (
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/application/rpc/client"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

type AuthService struct{}

func newAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) GenerateToken(dto *model.AuthRequest) (string, error) {
	var tokenString string
	if err := rpc_client.Call(config.AuthPort, "AuthService.Authenticate", &dto, &tokenString); err != nil {
		return "", err
	}
	return tokenString, nil
}
