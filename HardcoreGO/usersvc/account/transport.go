package account

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"github.com/gorilla/mux"
)

//  Structure for Request and Response
type CreateUserRequest struct{
	Email		string `json:"email"`
	Password	string `json:"password"`
}

type CreateUserResponse struct{
	Ok		string `json:"ok"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type GetUserResponse struct{
	Email string `json:"email"`
}


// Decoding Request from Transport

func DecodeCreateUserRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request CreateUserRequest

	if err:= json.NewDecoder(r.Body).Decode(&request); err!= nil{
		return nil, err
	}

	return request, nil
}

func DecodeGetUserRequest(ctx context.Context, r *http.Request)(interface{}, error){
	var request GetUserRequest

	vars:= mux.Vars(r)

	request = GetUserRequest{
		Id : vars["id"],
	}

	return request, nil
}


func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error{
	return json.NewEncoder(w).Encode(response)
}

// Making endpoint Adapter
type Endpoints struct {
	CreateUser 	endpoint.Endpoint
	GetUser		endpoint.Endpoint
}

func MakeEndpoints(s UserService) Endpoints{
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser   : makeGetUserEndpoint(s),
	}
}
func makeCreateUserEndpoint(s UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(CreateUserRequest)

		ok, err := s.CreateUser(ctx, req.Email, req.Password)

		return CreateUserResponse{ok}, err
	}
}

func makeGetUserEndpoint( s UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Email: email,
		}, err
	}
}