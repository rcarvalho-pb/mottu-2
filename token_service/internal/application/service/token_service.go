package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rcarvalho-pb/mottu-token_service/internal/application/config"
	"github.com/rcarvalho-pb/mottu-token_service/internal/model"
)

type TokenService struct {
	jwtSecret []byte
}

// NewTokenService cria uma nova instância do serviço de tokens
func NewTokenService() *TokenService {
	return &TokenService{
		jwtSecret: []byte(config.Secret),
	}
}

// GenerateJwt gera um token JWT para um usuário
func (s *TokenService) GenerateJwt(user *model.UserDTO) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := model.ClaimsDTO{
		UserId:   user.Id,
		Username: user.Username,
		UserRole: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "mottu-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar token: %w", err)
	}
	return tokenString, nil
}

// GetClaims extrai as claims de um token JWT
func (s *TokenService) ValidateToken(tokenString string) (*model.ClaimsDTO, error) {
	claims := &model.ClaimsDTO{}
	token, err := jwt.ParseWithClaims(tokenString, claims, s.getValidationKey)
	if err != nil {
		return nil, fmt.Errorf("erro ao analisar token: %w", err)
	}

	if claims, ok := token.Claims.(*model.ClaimsDTO); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token inválido ou claims ausentes")
}

// getValidationKey retorna a chave de validação do JWT
func (s *TokenService) getValidationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
	}

	return s.jwtSecret, nil
}
