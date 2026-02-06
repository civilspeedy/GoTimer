package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	tickerLength time.Duration = 1 * time.Second
	dateTemplate string        = "dd/mm/yyyy"
)

var (
	ticker    = time.NewTicker(tickerLength)
	seconds   uint
	stopChan  chan struct{}
	pauseChan = make(chan bool, 1)
	mu        sync.RWMutex
	paused    bool
)

const (
	secInHr  uint = 3600
	secInMin uint = 60
)

func secToStr(sec uint) string {
	mu.Lock()
	defer mu.Unlock()
	defer logTime()()

	hours := sec / secInHr
	minutes := (sec % secInHr) / secInMin
	theSeconds := sec % secInMin

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, theSeconds)
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
	mu.Lock()
	defer mu.Unlock()
	defer logTime()()

	paused = false
	stopChan = make(chan struct{})

	go tick()
}

func main() {
	start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
