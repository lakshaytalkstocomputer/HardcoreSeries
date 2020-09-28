package main

import (
	"fmt"
	"github.com/lakshaytalkstocomputer/HardcoreSeries/HardcoreGO/socialmedia"
	)

func main(){
	myPost := socialmedia.NewPost(
		"lakshaytalksto",
		socialmedia.Moods["thrilled"],
		"Go is awesome!",
		"Check out the website",
		"https://golang.org",
		"",
		"",
		[]string{"go", "golang", "programming language"},
		)

	fmt.Printf("My Post: %+v\n", myPost)

}
