package main

import (
	"errors"
	"fmt"
)

type msg = byte

const (
	invalid msg = iota
	noRun
	started
	stopped
	nowPaused
	alPause
	noPause
	resumed
	toStop
	saveQuest
)

func message(m msg) {
	defer logTime()()
	var toPrint string

	switch m {
	case invalid:
		toPrint = "Invalid input!"
	case noRun:
		toPrint = "Timer not running!"
	case started:
		toPrint = "Timer started."
	case stopped:
		toPrint = "Timer stopped."
	case nowPaused:
		toPrint = "Timer paused."
	case alPause:
		toPrint = "Timer already paused."
	case noPause:
		toPrint = "Timer not paused."
	case resumed:
		toPrint = "Timer resumed."
	case toStop:
		toPrint = "Timer is running, want to stop? (y/n)"
	case saveQuest:
		toPrint = "Would you like to save timer? (y/n)"
	default:
		t := trace(2)
		err := errors.New("No matching message!")
		errOut(err, t)
	}

	fmt.Println(toPrint)
}
