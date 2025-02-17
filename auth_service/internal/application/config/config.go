package config

import "os"

var (
	UserPort  string
	AuthPort  string
	TokenPort string
)

func Start() {
	AuthPort = os.Getenv("AUTH_SERVICE_PORT")
	UserPort = os.Getenv("USER_SERVICE_PORT")
	TokenPort = os.Getenv("TOKEN_SERVICE_PORT")
}
