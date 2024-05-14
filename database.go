package main

import (
	"database/sql"
	"time"
)

var POOL *sql.DB

func initDB() error {

	var err error

	POOL, err = sql.Open("mysql", "root:@tcp(db)/app")
	if err != nil {
		return err
	}

	POOL.SetConnMaxLifetime(time.Minute * 3)
	POOL.SetConnMaxIdleTime(time.Minute * 5)
	POOL.SetMaxIdleConns(1)
	POOL.SetMaxOpenConns(3)

	return nil

}
