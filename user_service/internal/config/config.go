package config

import (
	"log"
	"os"
)

type Config struct {
	DBPath string
	FSPath string
}

func Start() *Config {
	dbPath := os.Getenv("DB_PATH")
	fsPath := os.Getenv("IMAGES_DIRECTORY")
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
	return &Config{
		DBPath: dbPath,
		FSPath: fsPath,
	}
}
