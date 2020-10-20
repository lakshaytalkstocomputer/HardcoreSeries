package main

//func MyGreetHandler(w http.ResponseWriter, r *http.Request){
//	injection.Greet(w, "world")
//}

import (
	"github.com/lakshaytalkstocomputer/HardcoreSeries/HardcoreGO/tdd/httpserver"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{

}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}


func main(){
	//injection.Greet(os.Stdout, "Lakshay")
	//http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))

	//mocking.Countdown(os.Stdout, &mocking.ConfigurableSleeper{1* time.Second, time.Sleep})

	server := &httpserver.PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil{
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
