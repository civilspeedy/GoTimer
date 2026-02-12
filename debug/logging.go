package debugtools

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/fatih/color"
)

// Whether debug mode is enabled or not. out() & logTime() will not print if this is false.
var DebugMode = true

func CreateErr(err error) error {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		log.Fatalln("Unable to read stack")
		return nil
	}
	funcDetails := runtime.FuncForPC(pc)

	return fmt.Errorf("file:%s line:%d func:%s error:\n%s\n", file, line, funcDetails.Name(), err)
}

// If debug mode is enabled will print a function's name out in order to mark logic flow.
func MarkFunc() {
	if DebugMode {
		pc, _, _, ok := runtime.Caller(2)
		if !ok {
			err := errors.New("runtime caller failure")
			log.Fatalln(CreateErr(err))
		}
		funcDetails := runtime.FuncForPC(pc)
		color.Blue(funcDetails.Name())
	}
}
