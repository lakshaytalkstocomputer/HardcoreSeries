package sync

import "testing"

func TestCounter(t *testing.T){
	t.Run("incrmenting the counter 3 time leaves it at 3", func(t *testing.T) {

		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()
		expected := 3

		if got != expected{
			t.Errorf("got %d, expected %d", got, expected)
		}

	})
}