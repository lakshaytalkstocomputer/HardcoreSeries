package main

import (
	"fmt"
)

/*
	* We are going to model the years the Summer Olympics took place
       along with the city that held them, from 1984 onwards, with an
       enumerated list
*/

func main(){
	const (
		LosAngeles =  1984 + (iota * 4)
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted or will host the Summer Olympics in the provided year...")

	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-18s %-18v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18v \n", "Tokyo", Tokyo)
}