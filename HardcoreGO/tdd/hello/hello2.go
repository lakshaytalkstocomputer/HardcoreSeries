// In hello1.go, we wrote function first, to know how to declare functions,
//  But from now we will write test first.

// We Must have test of our function first in order to
//  iterate on our software safely.

// We should start by capturing the requirements in a test.
//  Head over to hello2_test.go to check the test and requirement.

// When you retrospectively write tests there is a risk that your
//  test may continue to pass even uf the code doesn't work as intended.
package main

// Constants should improve the performance of your application
//  as it saves you creating the "Hello, " string instance every time
//  Hello2() is called.
//
// Refactor the constants
const englishHelloPrefix = "Hello, "

// We write this function by listening to compiler
//  On First run, it tells that it did not find any function named Hello2()
//    Then we define a function named Hello2()
//  On Second run, it tells us that Hello2 should have an argument of type string
//    Then we change the function to accept string and return string
//  On third run, it tells us that the test that runs Hello2 fails,becuase it did not return the expected
//    Then we change the logic of function
func Hello2(name string) string{
	return englishHelloPrefix + name
}
