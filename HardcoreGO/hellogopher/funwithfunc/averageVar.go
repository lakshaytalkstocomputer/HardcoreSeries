package main


import (
	"fmt"
)

func getAverage(nums ... float32)float32{
	var sum float32 = 0.0
	for _ , value := range nums{
		sum += value
	}
	return sum / float32(len(nums))
}

func main(){
	fmt.Println("Average of 2,3,6,7,8 is :", getAverage(2,3,6,7,8))
}