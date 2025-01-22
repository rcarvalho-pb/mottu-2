package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/service"
)

type RPCServer struct {
	userService *service.UserService
	Port        string
}

func New(service *service.UserService, port string) *RPCServer {
	return &RPCServer{
		userService: service,
		Port:        port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.Port))
	if err != nil {
		return err
	}
	defer listen.Close()
	err = rpc.RegisterName("UserService", r)
	if err != nil {
		fmt.Println(err)
	}
	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			return err
		}
		go rpc.ServeConn(rpcConn)
	}
}

func (r *RPCServer) CreateUser(newUser *dto.UserDTO, _ *struct{}) error {
	if newUser == nil {
		return fmt.Errorf("user can't be null")
	}
	if err := r.userService.CreateUser(newUser); err != nil {
		return err
	}
	return nil
}

func (r *RPCServer) GetUserById(userId *int64, reply *dto.UserDTO) error {
	user, err := r.userService.GetUserById(*userId)
	if err != nil {
		return fmt.Errorf("error getting user by id: %s\n", err)
	}
	*reply = *user
	return err
}

func (r *RPCServer) GetUserByUsername(username *string, reply *dto.UserDTO) error {
	user, err := r.userService.GetUserByUsername(*username)
	if err != nil {
		return fmt.Errorf("error getting user by username: %s\n", err)
	}
	*reply = *user
	return err
}

func (r *RPCServer) GetAllActiveUsers(_ struct{}, reply *[]*dto.UserDTO) error {
	users, err := r.userService.GetAllActiveUsers()
	if err != nil {
		return fmt.Errorf("error getting all active users: %s\n", err)
	}
	*reply = users
	return err
}

func (r *RPCServer) GetAllUsers(_ struct{}, reply *[]*dto.UserDTO) error {
	users, err := r.userService.GetAllUsers()
	if err != nil {
		return fmt.Errorf("error getting all users: %s\n", err)
	}
	*reply = users
	return err
}

func (r *RPCServer) DeactivateUser(userId *int64, reply *bool) error {
	if err := r.userService.DeactivateUser(*userId); err != nil {
		return fmt.Errorf("error deactivating user [%d]: %s\n", userId, err)
	}
	*reply = true
	return nil
}

func (r *RPCServer) ReactivateUser(userId int64, _ *struct{}) error {
	if err := r.userService.ActivateUser(userId); err != nil {
		return fmt.Errorf("error reactivating user [%d]: %s\n", userId, err)
	}
	return nil
}

func (r *RPCServer) ValidatePassword(userDTO *dto.UserDTO, _ *struct{}) error {
	if err := r.userService.ValidatePassword(userDTO); err != nil {
		return fmt.Errorf("Passwords doesn't match")
	}
	return nil
}

func (r *RPCServer) UpdatePassord(userDTO *dto.UserDTO, _ *struct{}) error {
	if err := r.userService.UpdateUser(userDTO); err != nil {
		return err
	}
	return nil
}
