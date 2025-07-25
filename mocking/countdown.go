package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

const (
	countdownStart = 3
	finalWrod      = "Go!"
	gapTime        = 1 * time.Second

	sleep = "sleep"
	write = "write"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (spy *SpySleeper) Sleep() {
	spy.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (spy *SpyCountdownOperations) Sleep() {
	spy.Calls = append(spy.Calls, sleep)
}

func (spy *SpyCountdownOperations) Write(value []byte) (n int, err error) {
	spy.Calls = append(spy.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (sleeper *ConfigurableSleeper) Sleep() {
	sleeper.sleep(sleeper.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (spy *SpyTime) SetDurationSlept(duration time.Duration) {
	spy.durationSlept = duration
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration, sleep}
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := range countdownFrom(3) {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWrod)
}

func countdownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
