package main

import (
	"fmt"
	"time"
)

func secToStr(sec uint) string {
	defer logTime()()

	hours := sec / secInHr
	minutes := (sec % secInHr) / secInMin
	theSeconds := sec % secInMin

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, theSeconds)
}

func dateToStr(date int64) string {
	defer logTime()()

	dateTime := time.Unix(date, 0)
	return fmt.Sprintf("%d/%d/%d", dateTime.Day(), dateTime.Month(), dateTime.Year())
}
