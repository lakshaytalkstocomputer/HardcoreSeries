package main

import "testing"

func TestHello2(t *testing.T){
	// Here we define the requirement of hello2 Function
	got := Hello2("Lakshay")
	want := "Hello, Lakshay"

	if got != want{
		t.Errorf("got %q want %q", got , want)
	}
}
