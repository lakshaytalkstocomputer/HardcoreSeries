package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface{
	Sleep()
}

type ConfigurableSleeper struct{
	Duration time.Duration
	Sleeep    func(time.Duration)
}
func(c *ConfigurableSleeper)Sleep(){
	c.Sleeep(c.Duration)
}

func Countdown(buffer io.Writer, sleeper Sleeper){

	for i := countdownStart ; i>0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(buffer, finalWord)


}