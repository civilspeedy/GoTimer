package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	tickerLength         = 1 * time.Second
	startMessage         = "Timer started."
	pauseMessage         = "Timer paused."
	resumeMessage        = "Timer resumed."
	stopMessage          = "Timer stopped."
	alreadyMessage       = "Timer already started."
	noTimerMessage       = "No active timer."
	alreadyPausedMessage = "Time already paused."
	notPausedMessage     = "Timer not paused."
	pausedCantStart      = "Timer paused, enter 'stop' to clear current timer or 'resume' to continue current"
	invalid              = "Invalid command, enter 'help' to see commands."
	stillRunning         = "Timer still running"
	wantToStop           = "Do you want to stop? (y/n)"
	wantToSave           = "Do you want to save time? (y/n)"
	addingTime           = "Adding time to database"
	commandHelp          = `
	start - Starts timer
	stop - Stops timer & prints final time
	pause - Pauses timer
	resume - Resumes timer after pausing
	reveal - shows current time
	debug - toggle debug mode
	`
)

var (
	ticker       = time.NewTicker(tickerLength)
	timer        = MyTime{seconds: 0}
	previousTime uint
	stopChan     chan struct{}
	pauseChan    = make(chan bool, 1)
	mu           sync.RWMutex

	running = false
	paused  = false
)

func in() string {
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

func startTimer() {
	mu.Lock()
	if running {
		if paused {
			fmt.Println(pausedCantStart)
		} else {
			fmt.Println(alreadyMessage)
		}
		mu.Unlock()
		return
	}

	fmt.Println(startMessage)
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

func stopTimer() {
	mu.Lock()
	if !running {
		fmt.Println(noTimerMessage)
		mu.Unlock()
		return
	}

	close(stopChan)
	running = false
	paused = false
	mu.Unlock()

	fmt.Println(stopMessage)
	fmt.Println(timer.toString())
	previousTime = timer.seconds
	timer.reset()
}

func pauseTimer() {
	if !running {
		fmt.Println(noTimerMessage)
	} else if paused {
		fmt.Println(alreadyPausedMessage)
	}
	paused = true

	select {
	case pauseChan <- true:
		fmt.Println(pauseMessage)
	default:

	}

}

func resumeTimer() {
	if !running {
		fmt.Println(noTimerMessage)
	} else if !paused {
		fmt.Println(notPausedMessage)
	} else {
		pauseChan <- false
		paused = false
		fmt.Println(stopMessage)
	}
}

func revealTimer() {
	if !running {
		fmt.Println(noTimerMessage)
	} else {
		fmt.Println("Timer is at: ", timer.toString())
	}
}

func savePrompt() {
	fmt.Println(wantToSave)
	if in() == "y" {
		fmt.Println(addingTime)
		insertTime(previousTime)
	}
}

func saveTimer() {
	var message string
	var state string
	if running {
		message = stillRunning
		state = "running"
	} else if paused {
		message = pauseMessage
		state = "paused"
	} else if !paused && !running {
		savePrompt()
	}
	fmt.Println(message + " " + wantToStop)
	if in() == "y" {
		stopTimer()
		savePrompt()
	} else {
		fmt.Println("Timer remains ", state)
	}
}

func main() {
	printDebug = true

	connectDatabase()
	defer database.Close()
	createTable()

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
			fmt.Println(commandHelp)
		case "debug":
			printDebug = !printDebug
			fmt.Println("Debug mode", printDebug)
		case "save":
			saveTimer()
		default:
			fmt.Println(invalid)
		}
	}
}
