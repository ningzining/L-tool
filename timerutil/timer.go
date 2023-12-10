package timerutil

import (
	"time"
)

func RunDelayTask(f func(), delayTime time.Duration) {
	time.AfterFunc(delayTime, f)
}

func RunPeriodicTask(f func(), delayTime time.Duration) {
	timer := time.NewTicker(delayTime)
	defer timer.Stop()

	f()
	for range timer.C {
		f()
	}
}
