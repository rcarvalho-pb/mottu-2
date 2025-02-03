package service

import (
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/application/rpc/client"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

const service = "UserService"

type UserService struct {
}

func newUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(dto *model.UserDTO) error {
	err := rpc_client.Call(config.UserPort, "UserService.CreateUser", &dto, &struct{}{})
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByiD(id int64) (*model.UserDTO, error) {
	var userDTO *model.UserDTO
	if err := rpc_client.Call(config.UserPort, "UserService.GetUserById", &id, &userDTO); err != nil {
		return nil, err
	}
	return userDTO, nil
}

func (us *UserService) GetUserByUsername(username string) (*model.UserDTO, error) {
	var userDTO *model.UserDTO
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.GetUserByUsername", service), &username, &userDTO); err != nil {
		return nil, err
	}
	return userDTO, nil
}

func (us *UserService) GetAllActiveUsers() ([]*model.UserDTO, error) {
	var usersDTO []*model.UserDTO
	if err := rpc_client.Call(config.UserPort, "UserService.GetAllActiveUsers", &struct{}{}, &usersDTO); err != nil {
		return nil, err
	}
	return usersDTO, nil
}

func (us *UserService) GetAllUsers() ([]*model.UserDTO, error) {
	var usersDTO []*model.UserDTO
	if err := rpc_client.Call(config.UserPort, "UserService.GetAllUsers", &struct{}{}, &usersDTO); err != nil {
		return nil, err
	}
	return usersDTO, nil
}

func (us *UserService) DeactivateUser(id int64) error {
	if err := rpc_client.Call(config.UserPort, "UserService")
}

func (us *UserService) UpdateUser(dto *model.UserDTO) error {
	if err := rpc_client.Call(config.UserPort, "UserService.UpdateUser", &dto, &struct{}{}); err != nil {
		return err
	}
	return nil
}
