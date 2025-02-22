package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}
