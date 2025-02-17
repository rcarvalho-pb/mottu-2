package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/service"
)

type RPCServer struct {
	service *service.Service
	Port    string
}

func New(service *service.Service, port string) *RPCServer {
	return &RPCServer{
		service: service,
		Port:    port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.Port))
	if err != nil {
		return err
	}
	defer listen.Close()
	err = rpc.RegisterName("AuthService", r)
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

func (r *RPCServer) Authenticate(authRequest *dto.AuthRequest, reply *string) error {
	tokenString, err := r.service.AuthenticateUser(authRequest)
	if err != nil {
		return err
	}
	*reply = tokenString
	return nil
}
