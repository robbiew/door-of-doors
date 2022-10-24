package timeout

import (
	"fmt"
	"time"
)

// Ticker holds time info
type Ticker interface {
	Duration() time.Duration
	Tick()
	Stop()
}

// TickFunc calls ticker
type TickFunc func(d time.Duration)

type ticker struct {
	*time.Ticker
	d time.Duration
}

var (
	stopTime bool
	// Make sure we can return cursor location after timer updates
	Warning     bool = false
	UpdatedTime int
)

func (t *ticker) Tick()                   { <-t.C }
func (t *ticker) Duration() time.Duration { return t.d }

// NewTicker creation
func NewTicker(d time.Duration) Ticker {
	return &ticker{time.NewTicker(d), d}
}

// Countdown for timer
func Countdown(ticker Ticker, duration time.Duration) chan time.Duration {
	remainingCh := make(chan time.Duration, 1)
	go func(ticker Ticker, dur time.Duration, remainingCh chan time.Duration) {
		for remaining := duration; remaining >= 0; remaining -= ticker.Duration() {
			if stopTime {
				ticker.Stop()
			}
			if remaining < time.Minute*5 {
				Warning = true
			}
			remainingCh <- remaining
			ticker.Tick()
		}
		ticker.Stop()
		close(remainingCh)

		fmt.Printf(" TIME'S UP!")

	}(ticker, duration, remainingCh)
	return remainingCh
}

func timer(timeLeft time.Duration) {
	for d := range Countdown(NewTicker(time.Minute), time.Minute*timeLeft) {
		if stopTime {
			UpdatedTime = int(d.Minutes())

			break
		} else {
			UpdatedTime = int(d.Minutes())

		}

	}
}

func StartTimer(timeLeft int) {
	stopTime = false
	go timer(time.Duration(timeLeft))
}
