package main

import (
	"os"
	"time"

	"github.com/gitKashish/learn-go/mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
