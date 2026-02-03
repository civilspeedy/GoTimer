package main

import (
	"database/sql"
	"fmt"
	"time"
	"timer/sqlDebug"
	"timer/sqlTime"

	_ "github.com/mattn/go-sqlite3"
)

const (
	path   = "./database.db"
	driver = "sqlite3"
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
func connectDatabase() error {
	defer logTime("Connect database")()

	out("Connecting database")
	db, err := sql.Open(driver, path)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	out("Database connected")
	database = db
	return nil
}

func checkDebugMode() (bool, error) {
	defer logTime("Check degbug mode value")()
	out("Fetching stored debug mode value")
	defer out("Fetched stored debug mode value")

	_, err := database.Exec(sqlDebug.Create)
	if err != nil {
		return false, err
	}

	rows, err := database.Query(sqlDebug.Select)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if !rows.Next() {
		_, err = database.Exec(sqlDebug.Insert, printDebug)
		return printDebug, nil
	}

	var d bool
	for rows.Next() {
		err = rows.Scan(&d)
		if err != nil {
			return false, nil
		}
	}
	return d, nil
}

func updateDebugMode(val bool) error {
	defer logTime("Update debug mode")()
	out("Updating stored debug mode")
	defer out("Updated stored debug mode")

	_, err := database.Exec(sqlDebug.Update, val)
	if err != nil {
		return err
	}
	printDebug = val
	return nil
}

// Drops the table "time".
func dropTable() error {
	defer logTime("Drop table")()
	out("Dropping table")

	_, err := database.Exec(sqlTime.Drop)
	out("Table Dropped")
	return err
}

// Create the table "time"
func createTable() error {
	defer logTime("Create table")()
	out("Creating table")
	_, err := database.Exec(sqlTime.Create)
	out("Table created")

	return err
}

// Returns all entries in "time" as TimeEntry slice.
func getAllTimes() ([]TimeEntry, error) {
	defer logTime("Get all times")()
	out("Fecthing all times")

	rows, err := database.Query(sqlTime.SelectAll)
	if err != nil {
		return nil, err
	}

	var rowSlice []TimeEntry
	for rows.Next() {
		var r TimeEntry
		err = rows.Scan(&r.id, &r.date, &r.time)
		if err != nil {
			return nil, err
		}
		rowSlice = append(rowSlice, r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	out("Fetched all times")
	return rowSlice, nil
}

// Using a formatted date string the total seconds of a recorded time is returned for that specific date.
func selectTime(dateStr string) (uint, error) {
	defer logTime("Select time")()
	out("Fetching time")

	rows, err := database.Query(sqlTime.SelectWhere, dateStr)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("No matching entry!")
	}

	var row uint
	err = rows.Scan(&row)
	if err != nil {
		return row, err
	}

	out("Time fetched")

	return row, nil
}

// Either updates or inserts passed seconds value in database.
func insertTime(recordedTime uint) error {
	defer logTime("Insert time")()
	out("Inserting time")
	date := newMyDate(time.Now())
	dateStr := date.toString()

	previousTime, err := selectTime(dateStr)

	if err != nil {
		out("Inserting new record")
		_, err = database.Exec(sqlTime.Insert, dateStr, recordedTime)
		if err != nil {
			return err
		}
		out("Inserted new record")
	} else {
		out("Updating previous record")
		_, err = database.Exec(sqlTime.Update, dateStr, previousTime+recordedTime) // not working
		if err != nil {
			return err
		}
		out("Updated previous record")
	}
	return nil
}
