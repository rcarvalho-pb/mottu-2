package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/middleware"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/router"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
)

func main() {
	config.Start()
	mux := router.New()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", config.BrokerPort),
		Handler: mux,
	}
	tService := service.NewTokenService()
	middleware.Init(tService)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
