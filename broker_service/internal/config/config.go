package config

import "os"

var (
	Secret     string
	BrokerPort string
	UserPort   string
	TokenPort  string
	AuthPort   string
)

func Start() {
	Secret = os.Getenv("MY_SECRET")
	BrokerPort = os.Getenv("BROKER_SERVICE_PORT")
	UserPort = os.Getenv("USER_SERVICE_PORT")
	AuthPort = os.Getenv("AUTH_SERVICE_PORT")
	TokenPort = os.Getenv("TOKEN_SERVICE_PORT")
}
