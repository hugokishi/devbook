package routes

import (
	"net/http"
	"web/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:                "/",
		Method:             http.MethodGet,
		Handler:            controllers.LoadLoginPage,
		NeedAuthentication: false,
	},
	{
		URI:                "/login",
		Method:             http.MethodGet,
		Handler:            controllers.LoadLoginPage,
		NeedAuthentication: false,
	},
	{
		URI:                "/login",
		Method:             http.MethodPost,
		Handler:            controllers.AuthenticateUser,
		NeedAuthentication: false,
	},
}
