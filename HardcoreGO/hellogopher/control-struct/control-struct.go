package main

import (
	"fmt"
)

func main(){
	// Let us declare an enum to represent odd and even number
	const (
		evenNumber uint8 = iota
		oddNumber
	)

	// Let us initialize a variable that would store a string
	const myString = "Hello My Name Is Lakshay!"

	// Let us loop overs the string and print characters at odd index

	for index , char := range myString{
		if uint8(index) % 2 == oddNumber{
			fmt.Println(index, ":" , string(char))
		}

	}

}
