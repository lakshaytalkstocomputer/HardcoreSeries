package main

import (
	f "fmt"
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

	f.Println("These cities hosted or will host the Summer Olympics in the provided year...")

	f.Printf("%-18s %-18s \n", "City", "Year")
	f.Printf("%-18s %-18v \n", "Milano Cortina", Milanocortina)
	f.Printf("%-18s %-18v \n", "Beijing", Beijing)
	f.Printf("%-18s %-18v \n", "Sochi", Sochi)
}