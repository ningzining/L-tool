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
	printR()
	time.Sleep(time.Hour)
}

func printLog() {
	log.Println("log")
}

func printR() {
	log.Println("R")
}
