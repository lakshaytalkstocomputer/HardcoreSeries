package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"flag"
)

func main(){
	var numTimes int
	var err error
	var gopherName string

	flag.StringVar(&gopherName, "gname", "Lakshay", "Name the Gopher")
	flag.Parse()
	fmt.Println("Hello", gopherName)

	for index, value := range os.Args{
		if index == 0{
			fmt.Println("The name of program is : ", value)
			continue
		}

		if index == 1{
			numTimes , err = strconv.Atoi(value)

			if err != nil{
				log.Fatal(err)
			}
		}


		fmt.Println("Value : ", value)
	}

	fmt.Println("Number of times: ", numTimes)



}
