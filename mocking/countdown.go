package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalWrod      = "Go!"
	gapTime        = 1 * time.Second
)

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(gapTime)
	}
	fmt.Fprintf(writer, finalWrod)
}
