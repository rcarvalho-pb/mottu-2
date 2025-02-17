package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rcarvalho-pb/mottu-token_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-token_service/internal/model"
)

type RPCServer struct {
	*service.TokenService
	port string
}

func New(port string, service *service.TokenService) *RPCServer {
	return &RPCServer{
		TokenService: service,
		port:         port,
	}
}

func (r *RPCServer) RPCListen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", r.port))
	if err != nil {
		return err
	}
	defer listen.Close()
	err = rpc.RegisterName("TokenService", r)
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

func (r *RPCServer) GenerateToken(dto *model.UserDTO, reply *string) error {
	tokenString, err := r.GenerateJwt(dto)
	if err != nil {
		return err
	}
	*reply = tokenString
	return err
}

func (r *RPCServer) ValidateToken(tokenString string, reply *model.ClaimsDTO) error {
	claims, err := r.GetClaims(tokenString)
	if err != nil {
		return fmt.Errorf("error validating token: %s", err)
	}
	*reply = *claims
	return nil
}
