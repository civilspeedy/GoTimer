package main

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
)

var printDebug bool

func out(msg string) {
	if printDebug {
		color.Green("LOG:%s\n", msg)
	}
}

func checkErr(err error) {
	if err != nil {
		color.Red("Fatal Error:")
		log.Fatalln(err)
		debug.PrintStack()
	}
}

func logTime(label string) func() {
	if printDebug {
		start := time.Now()
		return func() {
			color.Cyan("TIME: %s in %.3fms\n", label, float64(time.Since(start).Nanoseconds())/1e6)
		}
	}
	return nil
}
