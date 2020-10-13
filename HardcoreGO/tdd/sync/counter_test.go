package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T){


	t.Run("incrmenting the counter 3 time leaves it at 3", func(t *testing.T) {

		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("It Runs Safely concurrently", func(t *testing.T) {
		// Defining expected output value, and number of goroutiones that will run
		wantedCount := 1000

		// Initializing the counter
		counter := Counter{}

		// Creating a Wait group, this will allow us to sync goroutines
		var wg sync.WaitGroup
		// Telling Wait Group how many goroutines will run,
		//   Each finished goroutine will decrement this value by 1
		wg.Add(wantedCount)

		for i :=0; i< wantedCount; i++{

			// Running Goroutine
			go func(w *sync.WaitGroup){
				// Incrememnting the counter
				counter.Inc()
				// Finished the goruotine and decrementing the value
				w.Done()
			}(&wg)
		}

		// We wait for all goroutines to finish that is value of wg to become zero
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int)  {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}