package service

import (
	"errors"
	"log"
	"net/http"
	"strings"

	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/application/rpc/client"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (ts *TokenService) ValidateToken(r *http.Request) (*model.Claims, error) {
	tokenString, err := extractToken(r)
	if err != nil {
		return nil, err
	}
	var claims model.Claims
	if err := rpc_client.Call(config.TokenPort, "TokenService.ValidateToken", &tokenString, &claims); err != nil {
		log.Println(config.TokenPort)
		log.Println("error token service call")
		return nil, err
	}
	return &claims, nil
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("cabeçalho Authorization inválido ou ausente")
	}

	return parts[1], nil
}
