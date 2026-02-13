package main

import (
	"fmt"
	d "timer/debug"
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
	added
	updated
)

var msgMap = map[msg]string{
	invalid:   "Invalid input!",
	noRun:     "Timer not running!",
	started:   "Timer started.",
	stopped:   "Timer stopped.",
	nowPaused: "Timer paused.",
	alPause:   "Timer already paused.",
	noPause:   "Timer not paused.",
	resumed:   "Timer is running, want to stop? (y/n)",
	toStop:    "Would you like to save timer? (y/n)",
	saveQuest: "Would you like to save timer? (y/n)",
	added:     "Added new entry.",
	updated:   "Updated existing entry.",
}

func message(m msg) {
	defer d.MarkFunc()
	fmt.Println(msgMap[m])
}
