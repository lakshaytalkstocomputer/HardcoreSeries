package main

import "testing"


// Writing a test is just like writing a function, with a few rules
//  * It needs to be in file with a name like xxx_test.go
//  * The test function start with the word Test
//  * The test function takes only one argument only t *testing.T
//
// t of type *Testing.T is out "hook" into testing framework
//  so we can do things like t.Fail() when out tests fails.
func TestHello(t *testing.T){
	got := Hello()
	want := "Hello, World"

	if got != want {
		// This will print out the message and fail the test.
		t.Errorf("got %q want %q", got, want)
	}
}