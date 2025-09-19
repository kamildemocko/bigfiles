package tools

import (
	"fmt"
	"time"
)

type Spinner struct {
	cycler *cycler
	delay  time.Duration
	stop   chan struct{}
}

func NewSpinner() *Spinner {
	return &Spinner{
		cycler: newCycler(),
		delay:  130 * time.Millisecond,
		stop:   make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	go func() {
		for {
			select {
			case <-s.stop:
				fmt.Printf("\r")
				return
			default:
				fmt.Printf("\r%s", s.cycler.Next())
				time.Sleep(s.delay)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	close(s.stop)
}
