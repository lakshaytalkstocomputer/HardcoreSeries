package main


import (
	"fmt"
)

func populateIntegerArray(tempArray [5]int){

	tempArray[0] = 100
	tempArray[1] = 200
	tempArray[3] = 300
	tempArray[4] = 400
}

func populateIntegerArrayWithReturnValue(tempArray [5]int) [5]int {

	tempArray[0] = 100
	tempArray[1] = 200
	tempArray[3] = 300
	tempArray[4] = 400

	return tempArray
}

func beatlesArrayExample(){
	beatlesArray := [4]string{
		"John",
		"Paul",
		"Ringo",
		"George"}


	fmt.Println("The name of third band memeer: ",beatlesArray[2])
	fmt.Println("Total numbeer of members in Beatles : ",len(beatlesArray))
}

func main(){
	/*
	    * Array has to be defined with length
	    * Array cannot grow and  shrink in size
	    * Default value of array would be zero-value of data type used
	    * Array is defined as value so when passing array to function,
	       it as passed as value not as reference
	    * We can define in argument of function of which length this function accepts array and datatype as well
	    *
	*/

	var myArray [5]int
	fmt.Printf("Contents of myArray: %v\n\n",myArray)


	populateIntegerArray(myArray)
	fmt.Printf("Contents of myArray: %v\n\n",myArray)

	myArray = populateIntegerArrayWithReturnValue(myArray)
	fmt.Printf("Contents of myArray: %v\n\n",myArray)

	fmt.Println("Length of myArray : ", len(myArray))

	beatlesArrayExample()

	var someArray []int

	// someArray[0] = 10
	// someArray = myArray
	fmt.Println("Some array : ", someArray)
	fmt.Println("Length of Some array: ", len(someArray))



}
