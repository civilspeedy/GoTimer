package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
	d "timer/debug"
)

const (
	tickerLength time.Duration = 1 * time.Second
	dateTemplate string        = "dd/mm/yyyy"
)

var date uint

func in() (string, error) {
	defer d.MarkFunc()
	scanner := bufio.NewReader(os.Stdin)
	for {
		input, err := scanner.ReadString('\n')
		if err != nil {
			return "", d.CreateErr(err)
		}
		inLen := len(input)
		if inLen == 0 || inLen > 32 {
			message(invalid)

		} else {
			val := strings.Trim(strings.ToLower(input), "\n")
			return val, nil
		}
	}
}

func save() error {
	defer d.MarkFunc()
	if running {
		message(toStop)

		input, err := in()
		if err != nil {
			return err
		}
		if input == "y" {
			stop()
		}
		return nil
	}
	message(saveQuest)

	input, err := in()
	if err != nil {
		return err
	}

	if input == "y" {
		fetchedPrevious, err := slct(date)
		if err != nil {
			return err
		}

		if fetchedPrevious == nil {
			err = insert()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	err := connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	date = uint(time.Now().Unix())

	for {
		input, err := in()
		if err != nil {
			log.Fatalln(err)
		}
		switch input {
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
