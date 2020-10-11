package racer

import (
	"net/http"
	"time"
)

func Racer(firstUrl string, secondUrl string) (fastUrl string){

	firstDuration  := measureResponseTime(firstUrl)
	secondDuration := measureResponseTime(secondUrl)

	if firstDuration < secondDuration{
		fastUrl = firstUrl
	}else {
		fastUrl = secondUrl
	}

	return
}

func measureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}