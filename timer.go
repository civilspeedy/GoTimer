package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ticker    = time.NewTicker(tickerLength)
	stopChan  chan struct{}
	pauseChan = make(chan bool, 1)
	mu        sync.RWMutex
	running   bool
	paused    bool
	seconds   uint
	previous  uint
)

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
				mu.Lock()
				seconds++
				mu.Unlock()
				fmt.Printf("\033[s\033[1;1H\033[KTimer: %ss\033[u", secToStr(seconds))
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

func start() {
	defer message(started)
	mu.Lock()
	defer mu.Unlock()

	paused = false
	running = true
	stopChan = make(chan struct{})

	go tick()
}

func stop() {
	defer message(stopped)
	mu.Lock()
	defer mu.Unlock()

	if !running {
		message(noRun)
		return
	}

	close(stopChan)

	previous = seconds
	seconds = 0
}

func pause() {
	if !running {
		message(noRun)
	} else if paused {
		fmt.Println(alPause)
	}

	mu.Lock()
	paused = true
	mu.Unlock()
	message(nowPaused)
	pauseChan <- true
}

func resume() {
	if !running {
		fmt.Println(noRun)
	} else if !paused {
		message(noPause)
	} else {
		pauseChan <- false
		paused = false
		message(resumed)
	}
}
