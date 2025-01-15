package main

import (
	"fmt"
	"os"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapter/db/sqlite"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	db := sqlite.GetDB(dbPath)
	users, err := db.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, u := range users {
			fmt.Printf("%+v\n", u)
		}
	}
}
