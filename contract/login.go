package contract

import (
	"github.com/brunvieira/lotus"
	"time"
)

type LoginData struct {
	Login string
	Password string
}

type SigninData struct {
	*LoginData
	Name string
	LastName string
	DOB time.Time
}

var (
	SigninRouteContract = lotus.RouteContract{
		Label: "SigninRoute",
		Description: "Signs an user in to the service. Returns a 400 in case of error",
		Method: lotus.POST,
		Path: "/signin",
		Data: SigninData{},
	}
	LoginRouteContract = lotus.RouteContract{
		Label:       "LoginRoute",
		Description: "Returns a JWT token for valid users. Returns a 400 in case of error",
		Method:      lotus.POST,
		Path:        "/login",
		Data:        LoginData{},
	}
	LogoutRouteContract = lotus.RouteContract{
		Label: "LogoutRoute",
		Description: "Deletes the JWT token and ends the user session",
		Method: lotus.POST,
		Path: "/logout",
	}
	LoginServiceContract = lotus.ServiceContract{
		Label:           "LoginService",
		Description:     "Service in charge of login operations",
		Host:            "localhost",
		Namespace:       "login",
		Port:            8080,
		RoutesContracts: []lotus.RouteContract{
			SigninRouteContract,
			LoginRouteContract,
			LogoutRouteContract,
		},
	}
)