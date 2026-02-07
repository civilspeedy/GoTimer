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
	default:
		t := trace(2)
		err := errors.New("No matching message!")
		errOut(err, t)
	}

	fmt.Println(toPrint)
}
