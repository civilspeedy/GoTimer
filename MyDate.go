package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MyDate struct {
	day   byte
	month byte
	year  uint16
}

func newMyDate(date time.Time) MyDate {
	return MyDate{
		day:   byte(date.Day()),
		month: byte(date.Month()),
		year:  uint16(date.Year()),
	}
}

func myDateFromString(str string) MyDate {
	vals := strings.Split(str, "/")

	day, err := strconv.Atoi(vals[0])
	checkErr(err)
	month, err := strconv.Atoi(vals[1])
	checkErr(err)
	year, err := strconv.Atoi(vals[2])
	checkErr(err)

	return MyDate{
		day:   byte(day),
		month: byte(month),
		year:  uint16(year),
	}
}

func (d *MyDate) toString() string {
	return fmt.Sprintf("%02d/%02d/%d", d.day, d.month, d.year)
}
