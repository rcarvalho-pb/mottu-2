package service_impl

import (
	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/dto"
	rpc_client "github.com/rcarvalho-pb/mottu-auth_service/internal/application/rpc/client"
	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/service"
)

type userServiceImpl struct {
	port string
}

func NewUserService(userServiceAddr string) service.UserService {
	return &userServiceImpl{
		port: userServiceAddr,
	}
}

func (us *userServiceImpl) GetUser(username string) (*dto.UserDTO, error) {
	var userDto *dto.UserDTO
	if err := rpc_client.Call(us.port, "UserService.GetUserByUsername", username, &userDto); err != nil {
		return nil, err
	}

	return userDto, nil
}

func (us *userServiceImpl) ValidatePassword(passwords *dto.ComparePasswordsDTO) error {
	if err := rpc_client.Call(us.port, "UserService.ComparePasswords", &passwords, &struct{}{}); err != nil {
		return err
	}

	return nil
}
