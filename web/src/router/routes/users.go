package routes

import (
	"net/http"
	"web/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                "/register",
		Method:             http.MethodGet,
		Handler:            controllers.LoadRegisterPage,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Handler:            controllers.CreateUser,
		NeedAuthentication: false,
	},
}
