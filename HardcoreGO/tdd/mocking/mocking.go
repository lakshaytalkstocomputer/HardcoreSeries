package mocking

import (
	"fmt"
	"io"
	"time"
)

func Countdown(buffer io.Writer){

	for i := 3; i>0; i-- {
		time.Sleep( 1 * time.Second)
		fmt.Fprintln(buffer, i)
	}
	time.Sleep( 1 * time.Second)
	fmt.Fprintln(buffer, "Go!")


}