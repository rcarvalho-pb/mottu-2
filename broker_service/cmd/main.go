package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/router"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/config"
)

func main() {
	config.Start()
	mux := router.New()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", config.BrokerPort),
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
