package main

// Imports Over here
import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"os"
)

func main(){
	// Creating an instance of logger
	logger := log.NewLogfmtLogger(os.Stderr)

	// Create a  service
	var svc StringService
	svc = stringService{}
	svc = loggingMiddleWare{logger: logger, next: svc}

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
	logger.Log(http.ListenAndServe(":8080", nil))

}








