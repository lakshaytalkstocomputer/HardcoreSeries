package main

import "fmt"

func main(){
	/*
		* Map is a collection of key value pairs.
		* Maps are references.
		* They are dynamic.
		* We can add and remove elements from the map as we require.
		* Data is not sorted in map
	 */

	var someDict map[int32]string = make(map[int32]string)

	someDict[1] = "One"
	someDict[2] = "Two"
	someDict[3] = "Threes"

	fmt.Println(someDict)

	for key, value := range someDict{
		fmt.Println("Key :", key, "Value: ", value)
	}
 }