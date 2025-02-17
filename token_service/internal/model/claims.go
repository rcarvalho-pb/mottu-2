package model

import "github.com/golang-jwt/jwt/v5"

type ClaimsDTO struct {
	UserId   int64
	Username string
	UserRole string
	jwt.RegisteredClaims
}
