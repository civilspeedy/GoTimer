package main

import (
	"fmt"
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
