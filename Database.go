package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

const (
	path   = "./database.db"
	driver = "sqlite"
)

const (
	create = `
	CREATE TABLE IF NOT EXISTS time(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT UNIQUE,
		time INTEGER
	);`
	drop            = "DROP TABLE IF EXISTS time;"
	selectWhereDate = "SELECT time FROM time WHERE date = ?;"
	insertWhereDate = "INSERT INTO time(date, time) VALUES(?, ?);"
	updateWhereDate = "UPDATE time SET time = ? WHERE date = ?;"
)

var database *sql.DB

func connectDatabase() {
	defer logTime("Connect database")()

	out("Connecting database")
	db, err := sql.Open(driver, path)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	out("Database connected")
	database = db
}

func dropTable() {
	defer logTime("Drop table")()
	out("Dropping table")

	_, err := database.Exec(drop)
	checkErr(err)
	out("Table Dropped")
}

func createTable() {
	defer logTime("Create table")()
	out("Creating table")
	_, err := database.Exec(create)
	checkErr(err)
	out("Table created")
}

func getAllTimes() string {
	defer logTime("Get all times")()
	out("Fecthing all times")

	return ""
}

func selectTime(dateStr string) (uint, error) {
	defer logTime("Select time")()
	out("Fetching time")

	rows, err := database.Query(selectWhereDate, dateStr)
	checkErr(err)

	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("No matching entry!")
	}

	var row uint
	err = rows.Scan(&row)
	checkErr(err)

	out("Time fetched")

	return row, nil
}

func insertTime(recordedTime uint) {
	defer logTime("Insert time")()
	out("Inserting time")
	date := newMyDate(time.Now())
	dateStr := date.toString()

	previousTime, err := selectTime(dateStr)

	if err != nil {
		out("Inserting new record")
		_, err = database.Exec(insertWhereDate, dateStr, recordedTime)
		checkErr(err)
		out("Inserted new record")
	} else {
		out("Updating previous record")
		_, err = database.Exec(updateWhereDate, dateStr, previousTime+recordedTime)
		checkErr(err)
		out("Updated previous record")
	}
}
