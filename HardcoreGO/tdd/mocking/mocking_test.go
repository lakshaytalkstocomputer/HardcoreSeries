package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func TestCountdown(t *testing.T){

	t.Run("print 3 2 1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		expected := `3
2
1
Go!
`
		if got != expected{
			t.Errorf("Got %s, excpectd %s", got , expected)
		}
	})

	t.Run("check number of calls to sleep()", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		if spySleeper.Calls != 4{
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}

	})

}