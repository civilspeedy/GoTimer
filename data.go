package main

import (
	"database/sql"
	"errors"
	dt "timer/debugTools"

	_ "github.com/mattn/go-sqlite3"
)

const (
	fileName = "./store.db"
	driver   = "sqlite3"
)

var (
	db *sql.DB
)

// Queiries
const (
	createTimersTable = `
	CREATE TABLE IF NOT EXISTS timers (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date INTEGER UNIQUE,
	timer INTEGER
	);
	`
	dropTimersTable = "DROP TABLE IF EXISTS timers;"
	insertTimers    = "INSERT INTO timers(date, timer) VALUES(?,?);"
	updateTimers    = "UPDATE timers SET = ? WHERE ? = ?;"
	selectTimers    = "SELECT * FROM timers WHERE ? = ?"
)

type Queiry = byte

const (
	create Queiry = iota
	drop
	insert
	update
	slct
)

func connect() *dt.TracableError {
	defer dt.LogTime()()

	db, err := sql.Open(driver, fileName)
	if err != nil {
		return dt.NewTE(err)
	}

	err = db.Ping()
	if err != nil {
		return dt.NewTE(err)
	}

	return nil
}

func execute(q Queiry, args []any) *dt.TracableError {
	defer dt.LogTime()()

	checkArgs := func() *dt.TracableError {
		if args != nil || len(args) != 0 {
			err := errors.New("Arguments such me nil or empty when creating table")
			return dt.NewTE(err)
		}
		return nil
	}

	switch q {
	case create:
		argCheck := checkArgs()
		if argCheck != nil {
			return argCheck
		}
		_, err := db.Exec(createTimersTable)
		if err != nil {
			return dt.NewTE(err)
		}
	case drop:
		argCheck := checkArgs()
		if argCheck != nil {
			return argCheck
		}
		_, err := db.Exec(dropTimersTable)
		if err != nil {
			return dt.NewTE(err)
		}
	case insert:
		if len(args) != 2 {
		}

	}
}
