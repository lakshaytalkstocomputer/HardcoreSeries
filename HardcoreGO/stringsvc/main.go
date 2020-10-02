package main

// Imports Over here
import (
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main(){
	// Create a  service
	svc := stringService{}

	// Create Handler for the functions
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
		)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
		)

	http.Handle("/uppercase",uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}








