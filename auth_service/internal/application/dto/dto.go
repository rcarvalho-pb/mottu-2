package dto

import "github.com/golang-jwt/jwt/v5"

type UserDTO struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ClaimsDTO struct {
	UserID   int64
	Username string
	UserRole string
	jwt.RegisteredClaims
}

type ComparePasswordsDTO struct {
	HashedPassword string
	Password       string
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
