package main


import "fmt"

// someVaraible := 10000
// // Above is an error as short hand is used to declare and initialize local variable
// // We need to use var

func main() {
	// Initializing variables simple

	var varName1 int                               // Declaring a int variable without initializing
	fmt.Printf("Value of varName1 %f\n", varName1) // This should print zero-value of int datatype
	fmt.Println("Value of varName1 :", varName1)

	var varName2 int = 27081997                    // Declaring and initializing an int variable
	fmt.Printf("Value of varName2 %f\n", varName2) // This should print the value that is initialized
	fmt.Println("Value of varName2 :", varName2)

	var varName3 string                            // Declaring a string variable without initializing
	fmt.Printf("Value of varName3 %f\n", varName3) // This should print zero-value fo string
	fmt.Println("Value of varName3 :", varName3)

	varName4 := 1234222222222                      // Short hand notation to declare and initialize a variable
	fmt.Printf("Value of varName4 %f\n", varName4) // This should print value assinged above
	fmt.Println("Value of varName4 :", varName4)

	var var1, var2, var3 = 145, "hello", "123.4"
	fmt.Println("Value of var1 :", var1)
	fmt.Println("Value of var2 :", var2)
	fmt.Println("Value of var3 :", var3)

	var4, var5, var6 := "Bye", 777, 1999.66
	fmt.Println("Value of var4 :", var4)
	fmt.Println("Value of var5 :", var5)
	fmt.Println("Value of var6 :", var6)

}