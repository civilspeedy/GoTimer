package debugtools

import (
	"time"

	"github.com/fatih/color"
)

// Whether debug mode is enabled or not. out() & logTime() will not print if this is false.
var DebugMode = true

func ErrOut(t *TracableError) {
	color.Red(
		"Error in file:%s func:%s line: %d [%s]\n",
		t.stack.fileName,
		t.stack.funcName,
		t.stack.line, t.err,
	)
}

// Messures and logs the execution time of a function in miliseconds.
func LogTime() func() {
	if DebugMode {
		stack := trace(2)

		start := time.Now()
		return func() {
			color.Cyan("%s:%.3fms\n", stack.funcName, float64(time.Since(start).Nanoseconds())/1e6)
		}
	}
	return func() {}
}
