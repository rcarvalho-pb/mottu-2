package service

import (
	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/dto"
)

type Service struct {
	us UserService
	ts TokenService
}

func New(userService UserService, tokenService TokenService) *Service {
	return &Service{userService, tokenService}
}

func (s *Service) AuthenticateUser(request *dto.AuthRequest) (string, error) {
	user, err := s.us.GetUser(request.Username)
	if err != nil {
		return "", err
	}
	passwords := &dto.ComparePasswordsDTO{
		HashedPassword: user.Password,
		Password:       request.Password,
	}
	if err = s.us.ValidatePassword(passwords); err != nil {
		return "", err
	}
	tokenString, err := s.ts.GetToken(user)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
