package service

import "github.com/rcarvalho-pb/mottu-auth_service/internal/application/dto"

type UserService interface {
	GetUser(string) (*dto.UserDTO, error)
	ValidatePassword(*dto.ComparePasswordsDTO) error
}

type TokenService interface {
	GetToken(*dto.UserDTO) (string, error)
}
