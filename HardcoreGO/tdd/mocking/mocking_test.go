package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}


const write = "write"
const sleep = "sleep"

type CountdownOperationsSpy struct{
	Calls []string
}

func(c *CountdownOperationsSpy) Sleep(){
	c.Calls = append(c.Calls, sleep)
}

func(c *CountdownOperationsSpy)Write(p []byte)(n int, err error){
	c.Calls = append(c.Calls, write)
	return
}

type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T){
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime{
		t.Errorf("should have slpet for %v but actually slept for %v", sleepTime, spyTime.durationSlept)
	}
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

	t.Run("check oreder of operations", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}

		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls){
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
}