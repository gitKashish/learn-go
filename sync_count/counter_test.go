package synccount

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		want := 3

		for range want {
			counter.Inc()
		}

		assertCounter(t, counter, want)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for range want {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, want)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	value := got.Value()

	if value != want {
		t.Errorf("got %d want %d", got, want)
	}
}
