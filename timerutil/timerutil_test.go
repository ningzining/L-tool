package timerutil

import (
	"log"
	"testing"
	"time"
)

func TestRunDelayTask(t *testing.T) {
	RunDelayTask(printLog, time.Second)
}

func TestRunPeriodicTask(t *testing.T) {
	RunPeriodicTask(printLog, time.Second)
}

func printLog() {
	log.Println("log")
}
