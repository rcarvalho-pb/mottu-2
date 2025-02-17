package service_impl

import (
	"fmt"
	"log"

	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/dto"
	rpc_client "github.com/rcarvalho-pb/mottu-auth_service/internal/application/rpc/client"
	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/service"
)

type tokenServiceImpl struct {
	addr string
}

func NewTokenService(tokenServiceAddr string) service.TokenService {
	return &tokenServiceImpl{
		addr: tokenServiceAddr,
	}
}

func (ts *tokenServiceImpl) GetToken(user *dto.UserDTO) (string, error) {
	var tokenString string
	if err := rpc_client.Call(ts.addr, "TokenService.GenerateToken", &user, &tokenString); err != nil {
		log.Println("here", err)
		return "", fmt.Errorf("error calling token service: %s", err)
	}

	return tokenString, nil
}
