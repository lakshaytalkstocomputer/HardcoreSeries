package main

import "fmt"

func main(){

	var subject string = "Gopher"

	fmt.Println("First Element of Gopher string : ",string("Gopher"[0]))
	fmt.Println("First Element of string stored in variable subject : ",string(subject[0]))
	fmt.Println("First Element of string stored in variable subject : ",string(subject[len(subject)-1]))
	fmt.Println("Hello "+subject+"!")
}