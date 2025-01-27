package main

import (
	"log"

	"github.com/rcarvalho-pb/mottu-user_service/internal/config"
)

func main() {
	conf := config.Start()
	log.Fatal(conf.RPCServer.RPCListen())
}
