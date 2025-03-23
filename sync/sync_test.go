package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup // * A WaitGroup waits for a collection of goroutines to finish
		wg.Add(wantedCount)   // * Set the number of goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // * Each go routine runs and calls Done when finished
			}()
		}
		wg.Wait() // * Block until all goroutines have finished

		assertCounter(t, counter, wantedCount)
	})
}

// * Initialize Counter as a pointer becuase it contains a Mutex.
// * A mutex must not be copied after first use
func NewCounter() *Counter {
	return &Counter{}
}
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
