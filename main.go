package main

import (
	"bufio"
	"os"
	"strings"
	"time"
	dt "timer/debugTools"
)

const (
	tickerLength time.Duration = 1 * time.Second
	dateTemplate string        = "dd/mm/yyyy"
)

var date int64

func in() string {
	scanner := bufio.NewReader(os.Stdin)
	for {
		input, err := scanner.ReadString('\n')
		if err != nil {
			dt.ErrOut(dt.NewTE(err))
		}
		inLen := len(input)
		if inLen == 0 || inLen > 32 {
			message(invalid)
		} else {
			val := strings.Trim(strings.ToLower(input), "\n")
			return val
		}
	}
}

func save(date int64) error {
	if running {
		message(toStop)
		if in() == "y" {
			stop()
		}
		return nil
	}

	message(saveQuest)
	if in() == "y" {

	}

	return nil
}

func main() {
	date = time.Now().Unix()

	for {
		switch in() {
		case "start":
			start()
		case "pause":
			pause()
		case "resume":
			resume()
		case "stop":
			stop()
		case "exit":
			os.Exit(0)
		}
	}
}
