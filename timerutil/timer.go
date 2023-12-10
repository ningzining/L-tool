package timerutil

import (
	"time"
)

func RunDelayTask(f func(), delayTime time.Duration) {
	time.AfterFunc(delayTime, f)
}

func RunPeriodicTask(f func(), d time.Duration) {
	timer := time.NewTicker(d)
	go f()
	go func() {
		for range timer.C {
			f()
		}
	}()
}
