package greetings

import "fmt"

func PrintGreeting(how string){
	if how == "Hi" || how == "hi"{
		printHi()
	} else {
		printBye()
	}

}

func printHi(){
	fmt.Println("Hi Gopher")
}

func printBye(){
	fmt.Println("Bye Gopher")
}

func PrintMagicNumber(){
	fmt.Println(GetValue())
}

func PrintMagicNumberDirect()  {
	fmt.Println(someMagicNumber)
}
