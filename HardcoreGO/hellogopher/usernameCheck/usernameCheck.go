package main

import (
	"fmt"
	"flag"
	"log"
	"regexp"
)

func main(){
	// Get username from command ine argument
	var usernameInput string
	flag.StringVar(&usernameInput, "username", "", "Username that you ant you check the validity of")
	flag.Parse()

	fmt.Println("GopherFace username validation Checker")
	fmt.Println("Checking Syntax for username :",
					usernameInput,
					"  resulted in: ",
					CheckUsernameSyntax(usernameInput),
					"\n")
}

func CheckUsernameSyntax(username string) bool {

	// REGEX FOR USERNAME
	const usernameRegex string = `^@?(\w){1,15}$`

	validationResult :=  false

	reg, err := regexp.Compile(usernameRegex)

	if err != nil{
		log.Fatal(err)
	}

	validationResult = reg.MatchString(username)
	return validationResult
}
