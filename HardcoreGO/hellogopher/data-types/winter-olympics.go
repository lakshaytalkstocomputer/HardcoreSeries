package main

import (
	"fmt"
)

/*
	We are going to model the years and cities for the winter olympics
 */

func main(){

	const (
		Sarajevo = 1984 + (iota * 4)
		Calgary
		Albertville
		Lillehammer
		Nagano
		SaltLakeCity
		Turin
		Vancouver
		Sochi
		Pyeongchang
		Beijing
		Milanocortina
	)

	fmt.Println("These cities hosted or will host the Summer Olympics in the provided year...")

	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18v \n", "Milano Cortina", Milanocortina)
	fmt.Printf("%-18s %-18v \n", "Beijing", Beijing)
	fmt.Printf("%-18s %-18v \n", "Sochi", Sochi)
}