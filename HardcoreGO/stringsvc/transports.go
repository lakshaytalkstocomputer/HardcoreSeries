package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

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

