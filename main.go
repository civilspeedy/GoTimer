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
	stopChan  chan struct{}
	pauseChan = make(chan bool, 1)
	mu        sync.RWMutex
	paused    bool
	seconds   uint
	date      int64
)

const (
	secInHr  uint = 3600
	secInMin uint = 60
)

func secToStr(sec uint) string {
	defer logTime()()

	hours := sec / secInHr
	minutes := (sec % secInHr) / secInMin
	theSeconds := sec % secInMin

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, theSeconds)
}

func dateToStr(date int64) string {
	defer logTime()()

	dateTime := time.Unix(date, 0)
	return fmt.Sprintf("%d/%d/%d", dateTime.Day(), dateTime.Month(), dateTime.Year())
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

	date = time.Now().Unix()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
