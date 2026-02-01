package main

import (
	"fmt"
	"sync"
)

const (
	secInHr  uint = 3600
	secInMin uint = 60
)

type MyTime struct {
	seconds uint
	mutex   sync.Mutex
}

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

func (t *MyTime) inc() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.seconds++
}

func (t *MyTime) reset() {
	defer logTime("Reset time")()
	out("Reseting time")
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.seconds = 0
	out("Reset time")
}
