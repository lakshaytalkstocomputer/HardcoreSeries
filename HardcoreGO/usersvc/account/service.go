package account

import (
	"context"
)

// This Service would represent the functionality of out applciation
type UserService interface{
	CreateUser(ctx context.Context, email string, password string)(string , error)
	GetUser(ctx context.Context, id string)(string, error)
}

