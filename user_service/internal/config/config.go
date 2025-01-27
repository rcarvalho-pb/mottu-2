package config

import (
	"log"
	"os"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapter/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/repository"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/service"
	rpc_server "github.com/rcarvalho-pb/mottu-user_service/internal/rpc/server"
)

type Config struct {
	DBPath    string
	FSPath    string
	Service   *service.UserService
	RPCServer *rpc_server.RPCServer
}

func Start() *Config {
	dbPath := os.Getenv("DB_PATH")
	fsPath := os.Getenv("IMAGES_DIRECTORY")
	port := os.Getenv("USER_SERVICE_PORT")
	info, err := os.Stat(fsPath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(fsPath, os.ModePerm); err != nil {
			log.Fatal("couldn't create fs directory")
		}
	} else if err != nil {
		log.Fatal(err)
	} else if !info.IsDir() {
		log.Fatalf("path [%s] exist but it isn't a folder", fsPath)
	}
	svc := getService(getDB(dbPath), fsPath)
	return &Config{
		DBPath:    dbPath,
		FSPath:    fsPath,
		Service:   svc,
		RPCServer: rpc_server.New(svc, port),
	}
}

func getDB(dbPath string) repository.UserRepository {
	return sqlite.GetDB(dbPath)
}

func getService(rep repository.UserRepository, imagesDirectory string) *service.UserService {
	return service.New(rep, imagesDirectory)
}
