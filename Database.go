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
	selectAll       = "SELECT * FROM time;"
	insertWhereDate = "INSERT INTO time(date, time) VALUES(?, ?);"
	updateWhereDate = "UPDATE time SET time = ? WHERE date = ?;"
)

// Struct for extracting all the values in an database time entry.
type TimeEntry struct {
	// Unsigned integer representing the unique id of the entry.
	id uint
	// Date value formatted as string.
	date string
	// The total seconds of the recorded time.
	time uint
}

// Package-wide variable for interaction with database.
var database *sql.DB

// Connects or creates database. Assigns values to package-wide database variable.
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

// Drops the table "time".
func dropTable() {
	defer logTime("Drop table")()
	out("Dropping table")

	_, err := database.Exec(drop)
	checkErr(err)
	out("Table Dropped")
}

// Create the table "time"
func createTable() {
	defer logTime("Create table")()
	out("Creating table")
	_, err := database.Exec(create)
	checkErr(err)
	out("Table created")
}

// Returns all entries in "time" as TimeEntry slice.
func getAllTimes() []TimeEntry {
	defer logTime("Get all times")()
	out("Fecthing all times")

	rows, err := database.Query(selectAll)
	checkErr(err)

	var rowSlice []TimeEntry
	for rows.Next() {
		var r TimeEntry
		rows.Scan(&r.id, &r.date, &r.time)
		rowSlice = append(rowSlice, r)
	}

	out("Fetched all times")
	return rowSlice
}

// Using a formatted date string the total seconds of a recorded time is returned for that specific date.
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

// Either updates or inserts passed seconds value in database.
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
