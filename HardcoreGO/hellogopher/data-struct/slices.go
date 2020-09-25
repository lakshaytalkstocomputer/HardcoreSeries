package main

import "fmt"

func printArray(array []int32){
	fmt.Println("Array : ", array)
	fmt.Println("Length of Slice: ", len(array))
	fmt.Println("Capacity of Slice: ", cap(array))
}

func populateArray(array []int32){
	array[0] = 111
	array[1] = 222
	array[2] = 333
	array[3] = 444
	array[4] = 555

}


func main(){
	/*
	    * Slices can grow and shrink
	    * Slices are references
	    * make() to create slices
	    * cap() Capacity function
	    * append() to put new element to slice
	    *
	 */
	var someSlice []int32 = make([]int32, 5)
	printArray(someSlice)

	someSlice[0] = 1
	printArray(someSlice)

	populateArray(someSlice)
	printArray(someSlice)

	someSlice = append(someSlice, 10)
	printArray(someSlice)

	someSlice = append(someSlice, 11)
	printArray(someSlice)
}