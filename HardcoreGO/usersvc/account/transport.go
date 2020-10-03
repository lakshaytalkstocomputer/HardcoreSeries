package account

import (
	"context"
	"encoding/json"
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
