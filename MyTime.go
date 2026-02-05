package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	secInHr  uint = 3600
	secInMin uint = 60
)

// Struct for storing time values using a single unsigned integer containing the total seconds.
type MyTime struct {
	seconds uint
	mu      sync.Mutex
}

// Calculates the hours, minutes & seconds then returns a formatted string of the time in a readable format.
func (t *MyTime) toString() string {
	defer logTime("Time to string")()
	out("Converting time to string")
	t.mu.Lock()
	defer t.mu.Unlock()

	hours := t.seconds / secInHr
	minutes := (t.seconds % secInHr) / secInMin
	seconds := t.seconds % 60

	out("Converted time to string")
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func (t *MyTime) fromString(str string) error {
	defer logTime("String to time")()
	out("Converting string to time")
	defer out("Converted string to time")

	t.mu.Lock()
	defer t.mu.Unlock()

	vals := strings.Split(str, ":")
	if len(vals) != 3 {
		return errors.New("Invalid time string format")
	}

	hours, err := strconv.Atoi(vals[0])
	if err != nil {
		return err
	}
	minutes, err := strconv.Atoi(vals[1])
	if err != nil {
		return err
	}
	seconds, err := strconv.Atoi(vals[2])
	if err != nil {
		return err
	}

	t.seconds = uint(hours)*secInHr + uint(minutes)*secInMin + uint(seconds)

	return nil
}

// Increments the seconds value.
func (t *MyTime) inc() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.seconds++
}

// Decrements the seconds value.
func (t *MyTime) dec() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.seconds--
}

// Returns seconds value
func (t *MyTime) getSec() uint {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.seconds
}

// Returns the seconds value to 0.
func (t *MyTime) reset() {
	defer logTime("Reset time")()
	out("Resetting time")
	t.mu.Lock()
	defer t.mu.Unlock()
	t.seconds = 0
	out("Reset time")
}
