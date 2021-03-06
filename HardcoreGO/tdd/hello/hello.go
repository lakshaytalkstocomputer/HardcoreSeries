// This is not a good practice as the fmt.Println() is a side effect
//  and string "Hello World" is our domain. So it is very important
//  to separate these two out!

/*
func main(){
	fmt.Println("Hello, World")
}
*/

// Above function should be written as :
//func Hello() string{
//
//return "Hello, World"
//}
//



// In above, we wrote function first, to know how to declare functions,
//  But from now we will write test first.

// We Must have test of our function first in order to
//  iterate on our software safely.

// We should start by capturing the requirements in a test.
//  Head over to hello_test.go to check the test and requirement.

// When you retrospectively write tests there is a risk that your
//  test may continue to pass even uf the code doesn't work as intended.
package main

import "fmt"

// Constants should improve the performance of your application
//  as it saves you creating the "Hello, " string instance every time
//  Hello2() is called.
//
// Refactor the constants
const spanish ="spanish"
const french ="french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix  = "Bonjour, "

// We write this function by listening to compiler
//  On First run, it tells that it did not find any function named Hello()
//    Then we define a function named Hello()
//  On Second run, it tells us that Hello should have an argument of type string
//    Then we change the function to accept string and return string
//  On third run, it tells us that the test that runs Hello fails,because it did not return the expected
//    Then we change the logic of function

func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name
}

// This function will check the language and return appropriate prefix
func greetingPrefix(language string)(prefix string){

	switch language {
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix

	}
	return
}
func main(){
	fmt.Println(Hello("Lakshay", ""))
}
