package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Struct for storing date values with unsigned integers.
type MyDate struct {
	day   byte
	month byte
	year  uint16
}

// Creates new date struct from a passed golang time struct.
func newMyDate(date time.Time) MyDate {
	return MyDate{
		day:   byte(date.Day()),
		month: byte(date.Month()),
		year:  uint16(date.Year()),
	}
}

// Converts string into new date struct.
func myDateFromString(str string) (*MyDate, error) {
	defer logTime("Date from string")()
	out("Converting date from string")
	vals := strings.Split(str, "/")
	if len(vals) != 3 {
		return nil, errors.New("Invalid format")
	}

	day, err := strconv.Atoi(vals[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(vals[1])
	if err != nil {
		return nil, err
	}
	year, err := strconv.Atoi(vals[2])
	if err != nil {
		return nil, err
	}

	out("Date converted")

	return &MyDate{
		day:   byte(day),
		month: byte(month),
		year:  uint16(year),
	}, nil
}

// Returns formatted string of date value.
func (d *MyDate) toString() string {
	return fmt.Sprintf("%02d/%02d/%d", d.day, d.month, d.year)
}
