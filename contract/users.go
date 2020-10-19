package contract

import (
	"github.com/brunvieira/lotus"
	"time"
)

type UserId struct {
	id string
}

type User struct {
	*UserId
	Name string
	LastName string
	DOB time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	CreateUserRouteContract = lotus.RouteContract{
		Label: "CreateUser",
		Description: "Registers an user to the service. Returns an 400 in case of error",
		Method: lotus.POST,
		Path: "/add",
		Data: User{},
	}
	DeleteUserRouteContract = lotus.RouteContract{
		Label: "DeleteUser",
		Description: "Removes an user from the service. Returns an 400 in case of error",
		Method: lotus.POST,
		Path: "/remove",
		Data: UserId{},
	}
	RetrieveUserRouteContract = lotus.RouteContract{
		Label: "RetrieveUser",
		Description: "Retrieves an user from the service. Returns an 400 in case of error",
		Method: lotus.GET,
		Path: "/get/:id",
		Data: UserId{},
	}
	UsersServiceContract = lotus.ServiceContract{
		Label: "UsersService",
		Description: "Service for registering, de-registering and retrieving users",
		Host: "localhost",
		Namespace: "users",
		Port: 8081,
		RoutesContracts: []lotus.RouteContract{
			CreateUserRouteContract,
			DeleteUserRouteContract,
			RetrieveUserRouteContract,
		},
	}
)
