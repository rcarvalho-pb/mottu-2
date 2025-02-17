package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rcarvalho-pb/mottu-token_service/internal/application/config"
	"github.com/rcarvalho-pb/mottu-token_service/internal/model"
)

var jwtSecret = []byte(config.Secret)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) GenerateJwt(user *model.UserDTO) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.ClaimsDTO{
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
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func (t *TokenService) GetClaims(tokenString string) (*model.ClaimsDTO, error) {
	claims := &model.ClaimsDTO{}
	token, err := jwt.ParseWithClaims(tokenString, claims, getValidationKey)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func getValidationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("invalid signature method")
	}
	return jwtSecret, nil
}
