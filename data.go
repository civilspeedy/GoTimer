package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	fileName = "./store.db"
	driver   = ""
)

var (
	db *sql.DB
)

func connect() *ErrStack {
	defer logTime()()

	db, err := sql.Open(driver, fileName)
	if err != nil {
		return &ErrStack{
			err:   err,
			stack: trace(2),
		}
	}

	err = db.Ping()
	if err != nil {
		return
	}
	return nil
}
