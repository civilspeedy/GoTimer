package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"timer/messages"
)

const (
	tickerLength time.Duration = 1 * time.Second
	dateTemplate string        = "dd/mm/yyyy"
)

var (
	ticker       = time.NewTicker(tickerLength)
	timer        = MyTime{seconds: 0}
	previousTime uint
	stopChan     chan struct{}
	pauseChan    = make(chan bool, 1)
	mu           sync.RWMutex

	// Whether the timer is actively running.
	running = false

	// Whether the timer is paused.
	paused = false

	countDown = false
)

// Scans text input and returns string value.
func in() (string, error) {
	defer logTime("Input")()
	out("Getting user input")
	defer out("Got user input")

	var input string
	var err error
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		if len(input) <= 32 {
			break
		} else {
			fmt.Println(messages.BigInput)
		}
	}

	fmt.Println(messages.ClearPrevious)
	input = strings.Trim(input, "\n")
	return strings.ToLower(input), nil
}

func tick() {
	ticker := time.NewTicker(tickerLength)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			mu.RLock()
			isPaused := paused
			mu.RUnlock()

			if !isPaused {
				if !countDown {
					timer.inc()
				} else {
					timer.dec()
				}
			}
		case p := <-pauseChan:
			mu.Lock()
			paused = p
			mu.Unlock()
		case <-stopChan:
			return
		}
	}
}

func startTimer() {
	mu.Lock()
	defer mu.Unlock()

	if running {
		if paused {
			fmt.Println(messages.PausedCantStart)
		} else {
			fmt.Println(messages.AlreadyStarted)
		}
		return
	}

	fmt.Println(messages.Start)
	running = true
	paused = false
	stopChan = make(chan struct{})

	go tick()
}

// Stops timer and closes channel.
func stopTimer() {
	mu.Lock()
	if !running {
		fmt.Println(messages.NoTimer)
		mu.Unlock()
		return
	}

	close(stopChan)
	running = false
	paused = false
	mu.Unlock()

	fmt.Println(messages.Stop)

	fmt.Println(timer.toString())
	previousTime = timer.getSec()
	timer.reset()
}

func pauseTimer() {
	if !running {
		fmt.Println(messages.NoTimer)
	} else if paused {
		fmt.Println(messages.AlreadyPaused)
	}
	mu.Lock()
	paused = true
	mu.Unlock()
	fmt.Println(messages.Pause)
	pauseChan <- true
}

func resumeTimer() {
	if !running {
		fmt.Println(messages.NoTimer)
	} else if !paused {
		fmt.Println(messages.NotPaused)
	} else {
		pauseChan <- false
		paused = false
		fmt.Println(messages.Resume)
	}
}

func revealTimer() {
	if !running {
		fmt.Println(messages.NoTimer)
	} else {
		fmt.Println("Timer is at: ", timer.toString())
	}
}

func savePrompt() error {
	fmt.Println(messages.WantToSave)
	input, err := in()
	if err != nil {
		return err
	}

	if input == "y" {
		fmt.Println(messages.AddTime)
		err := insertTime(previousTime)
		if err != nil {
			return err
		}
	}
	return nil
}

// Prompts users if they want to stop the timer (if its running) & save.
func saveTimer() error {
	var message string
	var state string
	if running {
		message = messages.StillRunning
		state = "running"
	} else if paused {
		message = messages.Pause
		state = "paused"
	} else if !paused && !running {
		err := savePrompt()
		if err != nil {
			return err
		}
	} else {
		fmt.Println(message + " " + messages.WantToStop)
		input, err := in()
		if err != nil {
			return err
		}
		if input == "y" {
			stopTimer()
			err := savePrompt()
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Timer remains ", state)
		}
	}
	return nil
}

func printALlTimes() error {
	times, err := getAllTimes()
	if err != nil {
		return err
	}
	for _, r := range times {
		rTime := MyTime{seconds: r.time}
		fmt.Printf("Date: %s Time: %s\n", r.date, rTime.toString())
	}
	return nil
}

func search() error {
	for {
		fmt.Printf("Enter date in %s format:\n", dateTemplate)
		searchDate, err := in()
		if err != nil {
			return err
		}

		_, err = myDateFromString(searchDate)
		if err != nil {
			fmt.Println(messages.InvalidDate)
		} else {
			fetchedTime, err := selectTime(searchDate)
			if err != nil {
				fmt.Printf("%s\n%s", messages.NoSearch, messages.AnotherDate)
				input, err := in()
				if err != nil {
					return err
				}
				if input != "y" {
					break
				}
			} else {
				fmt.Printf("%s: %d\n", searchDate, fetchedTime)
				break
			}
		}
	}

	return nil
}

func export() error {
	times, err := getAllTimes()
	if err != nil {
		return err
	}

	var records []TimeEntry
	for _, r := range times {
		records = append(records, r)
	}

	err = writeToCSV(records)
	return err
}

func startCountdown() error {

	fmt.Println(messages.Countdown)
	input, err := in()
	if err != nil {
		return err
	}

	var duration MyTime
	if strings.Contains(input, ":") {
		err := duration.fromString(input)
		if err != nil {
			return err
		}
	} else {
		convertedInput, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		duration.seconds = uint(convertedInput)
	}

	countDown = true

	startTimer()

	return nil
}

func main() {
	checkErr(connectDatabase())
	printDebug, err := checkDebugMode()
	checkErr(err)
	defer database.Close()
	checkErr(createTable())

	for {
		input, err := in()
		checkErr(err)
		switch input {
		case "start":
			startTimer()
		case "stop":
			stopTimer()
		case "pause":
			pauseTimer()
		case "resume":
			resumeTimer()
		case "reveal":
			revealTimer()
		case "help":
			fmt.Printf(messages.CommandHelp)
		case "debug":
			printDebug = !printDebug
			fmt.Println("Debug mode", printDebug)
			checkErr(updateDebugMode(printDebug))
		case "save":
			checkErr(saveTimer())
		case "times":
			checkErr(printALlTimes())
		case "search":
			search()
		case "export":
			checkErr(export())
		case "clear":
			fmt.Println(messages.ClearAll)
		case "countdown":
			checkErr(startCountdown())
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("'%s' %s\n", input, messages.Invalid)
		}
	}
}
