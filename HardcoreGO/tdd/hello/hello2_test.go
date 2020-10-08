package main

import "testing"

func TestHello2(t *testing.T){
	// Here we define the requirement of hello2 Function
	//  Each requirement would be in  sub-test,
	//  We can write sub-test using t

	//  We first define the function which compares got and want.
	//   This function helps us to avoid duplication and improves readability
	//   of our tests.
	//
	//  t.Helper() is needed to tell the test suite that this method is a helper.
	//    By doing this when it fails the line number reported will be in our function
	//    call rather than inside our test helper
	assertCorrectMessage := func(t *testing.T, got string, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got ,want)
		}
	}

	// We can define the requirement as one sub test
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello2("Lakshay")
		want := "Hello, Lakshay"
		assertCorrectMessage(t, got, want)
	})

	// We define other requirement as other subtest
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello2("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
