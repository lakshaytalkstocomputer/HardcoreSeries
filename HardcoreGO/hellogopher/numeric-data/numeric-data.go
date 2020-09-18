package main

import "fmt"

func main(){

	// Let us declare 32 bit integer without intializing it
	var myInteger int32

	// We expect to see a  value of 0 since the zero-value of an int32 is 0
	fmt.Println("Value of myInteger: ", myInteger)

	// Declare and initliaze 2 float4's and then show their sum
	var firstFloat float64 =  3.333
	var secondFloat float64 = 306.9696969

	fmt.Println("Sum of my floating point numbers : ", firstFloat + secondFloat)

	x,y,z := 0,1,2
	fmt.Printf("x: %d\ty: %d\t z:%d\t\n", x,y,z)

	// Example of Complex number
	complexNumber := 5 + 24i
	fmt.Println("Value of Complex Number: ", complexNumber)

	// Example of grouping constant declaration / initializations
	const (
		speedOfLight    = 299792458
		pi              = 3.14
		myFavNum        = 108
	)

	// Example of grouping variable declarations/initializations
	var (
		a int = 5555
		b     = 8 + 31
		c     = 2.7
	)

	fmt.Printf("a: %#v\tb: %#v\tc: %#v\n",a,b,c)
}