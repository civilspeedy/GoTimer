package main

import (
	"fmt"
	"os"
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
)

// Scans text input and returns string value.
func in() string {
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

// Starts timer as a seperate go routine.
func startTimer() {
	mu.Lock()
	if running {
		if paused {
			fmt.Println(messages.PausedCantStart)
		} else {
			fmt.Println(messages.AlreadyStarted)
		}
		mu.Unlock()
	} else {
		fmt.Println(messages.Start)

		running = true
		paused = false
		stopChan = make(chan struct{})
		mu.Unlock()

		go func() {
			ticker := time.NewTicker(tickerLength)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					mu.RLock()
					isPaused := paused
					mu.RUnlock()
					if !isPaused {
						timer.inc()
					}
				case p := <-pauseChan:
					mu.Lock()
					paused = p
					mu.Unlock()
				case <-stopChan:
					return
				}
			}
		}()
	}
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
	if in() == "y" {
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
	}
	fmt.Println(message + " " + messages.WantToStop)
	if in() == "y" {
		stopTimer()
		err := savePrompt()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Timer remains ", state)
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

func search() {
	for {
		fmt.Printf("Enter date in %s format:\n", dateTemplate)
		searchDate := in()
		_, err := myDateFromString(searchDate)

		if err != nil {
			fmt.Println(messages.InvalidDate)
		} else {
			fetchedTime, err := selectTime(searchDate)
			if err != nil {
				fmt.Printf("%s\n%s", messages.NoSearch, messages.AnotherDate)
				if in() != "y" {
					break
				}
			} else {
				fmt.Printf("%s: %d\n", searchDate, fetchedTime)
				break
			}
		}

	}
}

func main() {
	printDebug = true

	checkErr(connectDatabase())
	defer database.Close()
	checkErr(createTable())

	for {
		switch in() {
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
			fmt.Println(messages.CommandHelp)
		case "debug":
			printDebug = !printDebug
			fmt.Println("Debug mode", printDebug)
		case "save":
			checkErr(saveTimer())
		case "times":
			checkErr(printALlTimes())
		case "search":
			search()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(messages.Invalid)
		}
	}
}
