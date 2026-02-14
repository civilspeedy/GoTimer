package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	d "timer/debug"
)

const (
	secInHr  uint = 3600
	secInMin uint = 60
)

func secToStr(sec uint) string {
	defer d.MarkFunc()
	return fmt.Sprintf("%02d:%02d:%02d",
		sec/secInHr,            // hours
		(sec%secInHr)/secInMin, // minutes
		sec%secInMin,           // seconds
	)
}

func dateToStr(date int64) string {
	defer d.MarkFunc()

	dateTime := time.Unix(date, 0)
	return fmt.Sprintf("%d/%d/%d", dateTime.Day(), dateTime.Month(), dateTime.Year())
}

func strToDate(str string) (uint, error) {
	defer d.MarkFunc()

	strSlice := strings.Split(str, "/")

	if len(strSlice) != 3 {
		err := errors.New("invalid string format for date")
		return 0, d.CreateErr(err)
	}

	var dateVals []int

	for _, value := range strSlice {
		val, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}

		dateVals = append(dateVals, val)
	}

	month := time.Month(dateVals[2])
	dateValue := time.Date(dateVals[2], month, dateVals[0], 0, 0, 0, 0, nil)

	return uint(dateValue.Unix()), nil
}
