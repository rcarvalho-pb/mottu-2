package main

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-user_service/internal/adapter/db/sqlite"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-user_service/internal/config"
)

func main() {
	fmt.Println("Starting....")
	conf := config.Start()
	fmt.Printf("%+v\n", conf)
	db := sqlite.GetDB(conf.DBPath)
	service := service.New(*db)
	users, err := db.GetAllUsers()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, u := range users {
			fmt.Printf("%+v\n", u)
		}
	}
	fmt.Println("Ending...")
}
