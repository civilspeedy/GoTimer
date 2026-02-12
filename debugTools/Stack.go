package debugtools

import (
	"log"
	"runtime"
)

// Structure for detailing the file, function and lines to be related to a log/error
type Stack struct {
	fileName string
	funcName string
	line     int
}

// Retruns stucture containing function details. This is costly and should not be called outside of debug mode or errors.
func trace(skip int) *Stack {
	defer LogTime()()
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
