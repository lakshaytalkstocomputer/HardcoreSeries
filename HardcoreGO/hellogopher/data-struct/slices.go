package main

import "fmt"

func printSlice(array []int32){
	fmt.Println("Array : ", array)
	//fmt.Println("Length of Slice: ", len(array))
	//fmt.Println("Capacity of Slice: ", cap(array))
}

func populateArray(array []int32){
	array[0] = 111
	array[1] = 222
	array[2] = 333
	array[3] = 444
	array[4] = 555

}

func printArray(array [5]int32){
	fmt.Println("Array : ", array)
	fmt.Println("Length of Array: ", len(array))
	fmt.Println("Capacity of Array: ", cap(array))
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
	printSlice(someSlice)

	someSlice[0] = 1
	printSlice(someSlice)

	populateArray(someSlice)
	printSlice(someSlice)

	someSlice = append(someSlice, 10)
	printSlice(someSlice)

	someSlice = append(someSlice, 11)
	printSlice(someSlice)

	/* */
	var arrayOfSomeKind = [5]int32{1,2,3,4,5}
	printArray(arrayOfSomeKind)
	fmt.Println()
	sliceFromArray := arrayOfSomeKind[0:3]
	println(sliceFromArray)
}