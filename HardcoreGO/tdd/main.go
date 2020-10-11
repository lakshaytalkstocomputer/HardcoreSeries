package main

import (
	"github.com/lakshaytalkstocomputer/HardcoreSeries/HardcoreGO/tdd/mocking"
	"os"
)

//func MyGreetHandler(w http.ResponseWriter, r *http.Request){
//	injection.Greet(w, "world")
//}

func main(){
	//injection.Greet(os.Stdout, "Lakshay")
	//http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))

	mocking.Countdown(os.Stdout)

}
