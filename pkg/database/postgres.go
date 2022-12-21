package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_Host   = "localhost"
	DB_Port   = "6433"
	DB_User   = "nbfriends"
	DB_Pass   = "nbfriends"
	DB_DbName = "nbfriends"
)

func ConnectDB() (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_Host, DB_Port, DB_User, DB_Pass, DB_DbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
