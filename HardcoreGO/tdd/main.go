package main

import "net/http"

import (
	"github.com/lakshaytalkstocomputer/HardcoreSeries/HardcoreGO/tdd/injection"
)


func MyGreetHandler(w http.ResponseWriter, r *http.Request){
	injection.Greet(w, "world")
}
func main(){
	//injection.Greet(os.Stdout, "Lakshay")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
