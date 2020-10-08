package main

import "fmt"

// This is not a good practice as the fmt.Println() is a side effect
//  and string "Hello World" is our domain. So it is very important
//  to separate these two out!

/*
func main(){
	fmt.Println("Hello, World")
}
*/

// Above function should be written as :
func Hello() string{
	return "Hello, world"
}

func main(){
	fmt.Println(Hello())
}

