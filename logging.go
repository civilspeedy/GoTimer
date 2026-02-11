package main

import (
	"errors"
	"log"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// Whether debug mode is enabled or not. out() & logTime() will not print if this is false.
var debugMode = true

type Stack struct {
	fileName string
	funcName string
	line     int
}

// Retruns stucture containing function details. This is costly and should not be called outside of debug mode or errors.
func trace(skip int) *Stack {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		log.Fatalln("Unable to read stack")
		return nil
	}

	funcDetails := runtime.FuncForPC(pc)

	return &Stack{
		fileName: file,
		funcName: funcDetails.Name(),
		line:     line,
	}

}

// Basic logging output for debugging messages.
func out(msg string) {
	if debugMode {
		color.Green("LOG:%s\n", msg)
	}
}

func errOut(err error, stack *Stack) {
	log.Fatalf("Error in file:%s func:%s line: %d\n%s\n", stack.fileName, stack.funcName, stack.line, err)
}

// Messures and logs the execution time of a function in miliseconds.
func logTime() func() {
	if debugMode {
		stack := trace(2)

		start := time.Now()
		return func() {
			color.Cyan("%s:%.3fms\n", stack.funcName, float64(time.Since(start).Nanoseconds())/1e6)
		}
	}
	return func() {}
}

// Prints error if function called inside has a runtime longer that provided time in seconds.
func checkTime(maxTime float64) func() {
	start := time.Now()

	return func() {
		runDuration := float64(time.Since(start).Nanoseconds()) / 1e6
		if runDuration >= maxTime*1000 {
			t := trace(3)
			err := errors.New("Function runtime took too long!")
			errOut(err, t)
		}
	}
}

type ErrStack struct {
	err   error
	stack *Stack
}
