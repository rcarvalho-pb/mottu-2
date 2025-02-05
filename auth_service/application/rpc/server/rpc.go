package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"
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
