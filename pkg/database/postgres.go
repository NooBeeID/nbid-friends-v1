package database

import (
	"backend/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var (
		DB_Host   = config.GetString(config.DB_HOST)
		DB_Port   = config.GetString(config.DB_PORT)
		DB_User   = config.GetString(config.DB_USER)
		DB_Pass   = config.GetString(config.DB_PASS)
		DB_DbName = config.GetString(config.DB_NAME)
	)
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
