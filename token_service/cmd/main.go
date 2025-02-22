package main

import (
	"log"

	"github.com/rcarvalho-pb/mottu-token_service/internal/application/config"
	rpc_server "github.com/rcarvalho-pb/mottu-token_service/internal/application/rpc/server"
	"github.com/rcarvalho-pb/mottu-token_service/internal/application/service"
)

func main() {
	config.Start()
	tokenService := service.NewTokenService()
	rpc := rpc_server.New(config.TokenPort, tokenService)
	log.Fatal(rpc.RPCListen())
}
