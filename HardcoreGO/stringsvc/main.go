package main

import "strings"

// Imports Over here
import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

/*
	* In Go Kit , We model a service as an interface
	*   Interface defines the type of operations on string
 */

type StringService interface{
	Uppercase(string) (string, error)
	Count(string) int
}


// stringService is the concrete implementation of StringService
type stringService struct{}

// Creating a type of error
var ErrEmpty = errors.New("Empty String")

// Implementing UpperCase operation from the interface
func (stringService) Uppercase(s string) (string, error)  {

	// Check if given string is empty
	if s == "" {
		return "", ErrEmpty
	}

	// Return the string and null error
	return strings.ToUpper(s), nil
}

// Implementing Count operation from the interface
func (stringService) Count(s string) int {
	return len(s)
}

/* Every Method in our interface will be modeled
 *  as a remote procedure call. For each method, wwe define request and response
 *   structs capturing all the input and output parameters respectively
*/

// Creating a uppercase Request type
type uppercaseRequest struct {
	S string `json:"s"`
}

// Creating upperCase response type
type uppercaseResponse struct {
	V string `json:"V"`
	Err string `json:"err,omitempty"` //Errors don't JSON
}

// Creating count request
type countRequest struct{
	S string `json:"s"`
}

// Creating count response
type countResponse struct{
	V int `json:"v"`
}

/*
	* An endpoint represents a single RPC. That is single method\
	*  That is single method in our service interface.
	*  We will write adapters to convert each of our service's methods
	*  into an endpoint.

	* Each adapter takes a String Service, and returns an endpoint
	*  that corresponds to one of the methods.
 */

// Creating an adapter for UpperCase endpoint
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)

		v, err := svc.Uppercase(req.S)

		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}

		return uppercaseResponse{v, ""}, nil
	}
}

// Creating an adapter for Count endpoint
func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

/*
	* Exposing our Endpoints
 */

func decodeUppercaseRequest(_ context.Context, r *http.Request)(interface{}, error){
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}

	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request)(interface{}, error){
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

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








