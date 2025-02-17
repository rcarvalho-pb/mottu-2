package config

import (
	"os"
)

var (
	TokenPort string
	Secret    string
)

func Start() {
	TokenPort = os.Getenv("TOKEN_SERVICE_PORT")
	Secret = os.Getenv("MY_SECRET")
}
