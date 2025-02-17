package service

import (
	"fmt"
	"log"

	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/application/rpc/client"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

const USER_RESOURCE = "UserService"

type UserService struct{}

func newUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(dto *model.UserDTO) error {
	err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.CreateUser", USER_RESOURCE), &dto, &struct{}{})
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByiD(id int64) (*model.UserDTO, error) {
	var userDTO *model.UserDTO
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.GetUserById", USER_RESOURCE), &id, &userDTO); err != nil {
		return nil, err
	}
	return userDTO, nil
}

func (us *UserService) GetUserByUsername(username string) (*model.UserDTO, error) {
	var userDTO *model.UserDTO
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.GetUserByUsername", USER_RESOURCE), &username, &userDTO); err != nil {
		return nil, err
	}
	return userDTO, nil
}

func (us *UserService) GetAllActiveUsers() ([]*model.UserDTO, error) {
	var usersDTO []*model.UserDTO
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.GetAllActiveUsers", USER_RESOURCE), &struct{}{}, &usersDTO); err != nil {
		return nil, err
	}
	return usersDTO, nil
}

func (us *UserService) GetAllUsers() ([]*model.UserDTO, error) {
	var usersDTO []*model.UserDTO
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.GetAllUsers", USER_RESOURCE), &struct{}{}, &usersDTO); err != nil {
		return nil, err
	}
	return usersDTO, nil
}

func (us *UserService) DeactivateUser(id int64) error {
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.DeactivateUser", USER_RESOURCE), &id, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *UserService) ReactivateUser(id int64) error {
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.ReactivateUser", USER_RESOURCE), &id, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdatePassword(dto *model.UpdatePasswordDTO) error {
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.UpdatePassword", USER_RESOURCE), &dto, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(dto *model.UserDTO) error {
	log.Printf("%+v\n", dto)
	if err := rpc_client.Call(config.UserPort, fmt.Sprintf("%s.UpdateUser", USER_RESOURCE), &dto, &struct{}{}); err != nil {
		return err
	}
	return nil
}
