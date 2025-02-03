package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return errors.New("invalid token")
	}
	return nil
}

func GetClaims(r *http.Request) (*model.Claims, error) {
	tokenString := extractToken(r)
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, getValidationKey)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func getValidationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(config.Secret), nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
