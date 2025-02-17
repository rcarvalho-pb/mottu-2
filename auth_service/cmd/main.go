package main

import (
	"log"

	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/config"
	rpc_server "github.com/rcarvalho-pb/mottu-auth_service/internal/application/rpc/server"
	"github.com/rcarvalho-pb/mottu-auth_service/internal/application/service"
	service_impl "github.com/rcarvalho-pb/mottu-auth_service/internal/application/service/impl"
)

func main() {
	config.Start()

	tokenSrv := service_impl.NewTokenService(config.TokenPort)
	userSrv := service_impl.NewUserService(config.UserPort)

	srv := service.New(userSrv, tokenSrv)

	rpc := rpc_server.New(srv, config.AuthPort)

	log.Fatal(rpc.RPCListen())
}
