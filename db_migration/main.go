package main

import (
	"database/sql"
	"embed"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed files/migrations/*.sql
var migrations embed.FS

func main() {
	path := os.Getenv("DB_PATH")
	db := Open(path)
	defer db.Close()

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "files/migrations"); err != nil {
		panic(err)
	}
}

func Open(connString string) *sql.DB {

	db := ConnectToDB(connString)

	if db == nil {
		panic("couldn't connect to db")
	}

	return db
}

func ConnectToDB(connString string) *sql.DB {
	count := 0

	for count < 10 {
		db, err := OpenDB(connString)
		if err == nil {
			return db
		}

		time.Sleep(1 * time.Second)
		count++
	}

	return nil
}

func OpenDB(connString string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
