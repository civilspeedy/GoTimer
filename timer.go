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
				fmt.Println(seconds)
				mu.Unlock()
			}
		case p := <-pauseChan:
			mu.Lock()
			paused = !p
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
	defer logTime()()

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

	seconds = 0
}
