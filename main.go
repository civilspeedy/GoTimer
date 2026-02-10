package main

import (
	"bufio"
	"os"
	"strings"
	"time"
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
			tr := trace(2)
			errOut(err, tr)
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
