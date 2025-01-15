package sqlite

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// GetUserById(int64) (*User, error)
// GetUserByUsername(string) ([]*User, error)
// CreateUser(User) error
// UpdateUser(User) error

var dbTimeout = 10 * time.Second

type DB struct {
	DB *sqlx.DB
}

func GetDB(dbPath string) *DB {
	db := connectToDB(dbPath)
	if db == nil {
		log.Fatal("couldn't connect to DB")
	}

	return &DB{db}
}

func connectToDB(dbPath string) *sqlx.DB {
	count := 0
	for count < 10 {
		db, err := openDB(dbPath)
		if err == nil {
			return db
		}

		count++
		time.Sleep(1 * time.Second)
	}

	return nil
}

func openDB(dbPath string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
