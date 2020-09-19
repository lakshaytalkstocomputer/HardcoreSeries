package main

import (
	"fmt"
)

func averageThree(firstNum float32, secondNum float32, thirdNum float32) float32 {
	var average = (firstNum + secondNum + thirdNum) / 3.0
	return average
}

func main(){
	fmt.Println("Average of 2,2,2 is : ", averageThree(2,2,2))
}