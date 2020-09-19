package greetings

import "fmt"

var someMagicNumber int

func init(){
	fmt.Println("Initializing someMagic Number")
	someMagicNumber = 27081997
}

func GetValue() int{
	return someMagicNumber
}