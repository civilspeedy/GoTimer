package main

import (
	"database/sql"
	"log"
	d "timer/debug"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func connect() error {
	defer d.MarkFunc()

	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		return d.CreateErr(err)
	}

	err = db.Ping()
	if err != nil {
		return d.CreateErr(err)
	}
	return nil
}

func closeDB() {
	err := db.Ping()
	if err != nil {
		log.Fatalln(d.CreateErr(err))
		return
	}

	err = db.Close()
	if err != nil {
		log.Fatalln(d.CreateErr(err))
	}
}

func create() error {
	defer d.MarkFunc()

	err := db.Ping()
	if err != nil {
		return d.CreateErr(err)
	}

	const quiery = `
	CREATE TABLE IF NOT EXISTS timers (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date INTEGER UNIQUE,
	seconds INTEGER
	);
	`
	_, err = db.Exec(quiery)
	if err != nil {
		return d.CreateErr(err)
	}
	return nil
}

func drop() error {
	defer d.MarkFunc()

	err := db.Ping()
	if err != nil {
		return d.CreateErr(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS timers;")
	if err != nil {
		return d.CreateErr(err)
	}

	return nil
}

func insert() error {
	defer d.MarkFunc()

	err := db.Ping()
	if err != nil {
		return d.CreateErr(err)
	}

	_, err = db.Exec("INSERT INTO timers(date, seconds) VALUES(?,?);", date, seconds)
	if err != nil {
		return d.CreateErr(err)
	}

	return nil
}

// Updates existing entries. Target is a boolean of etiher 'date' or 'seconds'.
func update(target bool, targetValue uint, newValue uint) error {
	defer d.MarkFunc()

	err := db.Ping()
	if err != nil {
		return d.CreateErr(err)
	}

	var collumMap = map[bool]string{
		true:  "date",
		false: "seconds",
	}

	_, err = db.Exec("UPDATE timers SET ? = ? WHERE ? = ?;",
		collumMap[target],
		newValue,
		collumMap[!target],
		targetValue,
	)
	if err != nil {
		return d.CreateErr(err)
	}

	return nil
}

func slct(date uint) (*uint, error) {
	defer d.MarkFunc()

	err := db.Ping()
	if err != nil {
		return nil, d.CreateErr(err)
	}

	rows, err := db.Query("SELECT seconds FROM timers WHERE date = ?", date)
	if err != nil {
		return nil, err
	}
	var sec uint
	if !rows.Next() {
		return nil, nil
	}

	for rows.Next() {
		err = rows.Scan(&sec)
		if err != nil {
			return nil, d.CreateErr(err)
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, d.CreateErr(err)
	}

	return &sec, nil
}
