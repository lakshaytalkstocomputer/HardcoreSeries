package account

import (
	"context"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler{
	r := mux.NewRouter()

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		DecodeCreateUserRequest,
		EncodeResponse,
		))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		DecodeGetUserRequest,
		EncodeResponse,
		))

	return r
}



