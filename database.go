package main

import (
	"database/sql"
	"time"
)

var pool *sql.DB

func initDB() error {

	pool, err := sql.Open("mysql", "root:@/app")
	if err != nil {
		return err
	}

	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetConnMaxIdleTime(time.Minute * 5)
	pool.SetMaxIdleConns(1)
	pool.SetMaxOpenConns(3)

	return nil

}
