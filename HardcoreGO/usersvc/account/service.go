package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// This Service would represent the functionality of out application
type UserService interface{
	CreateUser(ctx context.Context, email string, password string)(string , error)
	GetUser(ctx context.Context, id string)(string, error)
}

// Implementation of Interface UserService
type userService struct {
	repository Repository
	logger log.Logger
}

// Implements Business Logic of CreateUser
func (u userService) CreateUser(ctx context.Context, email string, password string) (string, error) {
	// Setting up Logger so that we know which method was called and with that
	logger := log.With(u.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id 		:= uuid.String()

	// Create instance of User
	user := User{
		ID : id,
		Email : email,
		Password: password,
	}

	// Use CreateUser method of repository to Persist the Data
	if err := u.repository.CreateUser(ctx, user); err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Created User", id)
	return "Success", nil
}

// Implements Business Logic of GetUser
func (u userService) GetUser(ctx context.Context, id string) (string, error) {
	// Setting up Logger  and logging to tell this method was called
	logger := log.With(u.logger, "method" , "GetUser")

	email, err := u.repository.GetUser(ctx, id)

	if err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return email, nil
}

// Creating A Service
func NewService(rep Repository, logger log.Logger) UserService {
	return &userService{
		repository: rep,
		logger: logger,
	}
}