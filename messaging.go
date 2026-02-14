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
	deletePrompt
	sureDeleteAll
	deletedAll
	deleteFailed
)

var msgMap = map[msg]string{
	invalid:       "Invalid input!",
	noRun:         "Timer not running!",
	started:       "Timer started.",
	stopped:       "Timer stopped.",
	nowPaused:     "Timer paused.",
	alPause:       "Timer already paused.",
	noPause:       "Timer not paused.",
	resumed:       "Timer is running, want to stop? (y/n)",
	toStop:        "Would you like to save timer? (y/n)",
	saveQuest:     "Would you like to save timer? (y/n)",
	added:         "Added new entry.",
	updated:       "Updated existing entry.",
	deletePrompt:  "Enter a specific date to delete entry or 'all' to remove all entries:",
	sureDeleteAll: "Are you sure? (y/n)",
	deletedAll:    "All entries are deleted.",
	deleteFailed:  "Failed to delete provided date, please try again.",
}

func message(m msg) {
	defer d.MarkFunc()
	fmt.Println(msgMap[m])
}
