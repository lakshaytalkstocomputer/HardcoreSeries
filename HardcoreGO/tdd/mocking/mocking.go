package mocking

import (
	"fmt"
	"io"
)

func Countdown(buffer io.Writer){

	for i := 3; i>0; i-- {
		fmt.Fprintln(buffer, i)
	}

	fmt.Fprintln(buffer, "Go!")


}