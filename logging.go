package main

import (
	"log"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
)

// Whether debug mode is enabled or not. out() & logTime() will not print if this is false.
var printDebug = false

// Basic logging output for debugging messages.
func out(msg string) {
	if printDebug {
		color.Green("LOG:%s\n", msg)
	}
}

// Checks if an error is nil and throws fatal if not. Will always be called regardless of debuging mode on or off.
func checkErr(err error) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		color.Red("Error[%s:%d]:", filename, line)
		log.Panic(err)
		debug.PrintStack()
	}
}

// Messures and logs the execution time of a function in miliseconds.
func logTime(label string) func() {
	if printDebug {
		start := time.Now()
		return func() {
			color.Cyan("TIME: %s in %.3fms\n", label, float64(time.Since(start).Nanoseconds())/1e6)
		}
	}
	return func() {}
}
