package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(firstUrl string, secondUrl string) (fastUrl string, err error){
	return ConfigurableRacer(firstUrl, secondUrl, tenSecondTimeout)
}

func ConfigurableRacer(firstUrl, secondUrl string, timeout time.Duration)(winner string, err error){
	select {

		case <- ping(firstUrl):
			return firstUrl, nil

		case <- ping(secondUrl):
			return secondUrl, nil

		case <- time.After(timeout):
			return "", fmt.Errorf("timed out waiting for %s and %s", firstUrl, secondUrl)
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