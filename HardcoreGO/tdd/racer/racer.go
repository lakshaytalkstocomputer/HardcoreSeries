package racer

import (
	"net/http"
)

func Racer(firstUrl string, secondUrl string) (fastUrl string){
	select {
		case <- ping(firstUrl):
			return firstUrl
		case <- ping(secondUrl):
			return secondUrl
	}
}

func ping(url string) chan struct{}{
	ch := make(chan struct{})

	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}