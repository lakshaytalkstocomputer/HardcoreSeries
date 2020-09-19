package main

import "fmt"

func multiPurpose(firstNumber int, secNumber int) (sum int, diff int){
	return firstNumber + secNumber, firstNumber  - secNumber
}

func main(){
	a, b := multiPurpose(10, 12)
	fmt.Println("A: ", a, "B: ", b)

	var alpha = func(){
		fmt.Println("Anonymous function")
	}

	fmt.Printf("%t\n", alpha)
	fmt.Println("Calling anon function below this!")

	alpha()
}
