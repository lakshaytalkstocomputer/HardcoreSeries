package main

import (
	"errors"
	"strings"
)

/*
* In Go Kit , We model a service as an interface
*   Interface defines the type of operations on string
 */

type StringService interface{
	Uppercase(string) (string, error)
	Count(string) int
}


// stringService is the concrete implementation of StringService
type stringService struct{}

// Creating a type of error
var ErrEmpty = errors.New("empty string")

// Implementing UpperCase operation from the interface
func (stringService) Uppercase(s string) (string, error)  {

	// Check if given string is empty
	if s == "" {
		return "", ErrEmpty
	}

	// Return the string and null error
	return strings.ToUpper(s), nil
}

// Implementing Count operation from the interface
func (stringService) Count(s string) int {
	return len(s)
}
