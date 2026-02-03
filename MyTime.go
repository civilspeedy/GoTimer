package main

import (
	"fmt"
	"sync"
)

const (
	secInHr  uint = 3600
	secInMin uint = 60
)

// Struct for storing time values using a single unsigned integer containing the total seconds.
type MyTime struct {
	seconds uint
	mutex   sync.Mutex
}

// Calculates the hours, minuts & seconds then returns a formatted string of the time in a readable format.
func (t *MyTime) toString() string {
	defer logTime("Time to string")()
	out("Converting time to string")
	t.mutex.Lock()
	defer t.mutex.Unlock()

	hours := t.seconds / secInHr
	minutes := (t.seconds % secInHr) / secInMin
	seconds := t.seconds % 60

	out("Converted time to string")
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// Increments the seconds value.
func (t *MyTime) inc() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.seconds++
}

// Returns seconds value
func (t *MyTime) getSec() uint {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.seconds
}

// Returns the seconds value to 0.
func (t *MyTime) reset() {
	defer logTime("Reset time")()
	out("Reseting time")
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.seconds = 0
	out("Reset time")
}
