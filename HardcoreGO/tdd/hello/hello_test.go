package main

import "testing"

func TestHello(t *testing.T){
	// Here we define the requirement of hello2 Function
	//  Each requirement would be in  sub-test, We can write sub-test using t.Run.
	//  It is useful to group tests around a "thing" and then have subtests describing
	//  different scenarios
	// Refactoring is not just for production code.
	// It is important that your tests are clear specifications of what the code needs to do.
	// We can and should refactor the tests.


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
		got := Hello("Lakshay", "")
		want := "Hello, Lakshay"
		assertCorrectMessage(t, got, want)
	})

	// We define other requirement as other subtest
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// We  now need to support a second parameter, specifying the language of a greeting.
	//  If a language that we do not need recognise, just default to english
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Lakshay", "spanish")
		want:= "Hola, Lakshay"
		assertCorrectMessage(t, got, want)
	})

	// Add support for French
	t.Run("in french", func(t *testing.T) {
		got := Hello("Lakshay", "french")
		want := "Bonjour, Lakshay"

		assertCorrectMessage(t, got, want)
	})

	// Test Unknown language default to engloish
	t.Run("in unknown language", func(t *testing.T) {
		got  := Hello("Lakshay", "Klingon")
		want := "Hello, Lakshay"

		assertCorrectMessage(t, got, want)
	})
}
