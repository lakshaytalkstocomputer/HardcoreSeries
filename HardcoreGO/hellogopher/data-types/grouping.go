package main

import (
	"fmt"
	"math"
)

func main(){
	// Here we group different variables to form a enum
	var (
			a int  = 1
			b      = 2
			c      = 3
	)
	fmt.Println(" A: ",a, " B:", b , "C: ", c)

	// Here wwe try to group together constants
	const (
		speedOfLight uint = 299792458
		pi              = 3.14
		myFavNum        = 108
	)
	fmt.Println("Speed Of Light: ", speedOfLight,
		              "Pi:", pi,
		  			  "MyFavNum: ", myFavNum)

	// Let us use generator to fill values to constant in enum
	const (
		alpha = iota
		beta
		gamma
		delta
	)
	fmt.Println("Vallue of Alpha: ", alpha)
	fmt.Println("Vallue of beta: ", beta)
	fmt.Println("Vallue of gamma: ", gamma)
	fmt.Println("Vallue of delta: ", delta)

	// Let ius use generator to fiull value to variables in enum
	// // Below code is an error as we require const to use iota
	//var (
	//	i = iota
	//	k
	//	l
	//
	//)

	fmt.Println("Value of log 2 in math lib : ", math.Ln2)


}