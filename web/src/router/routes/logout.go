package routes

import (
	"net/http"
	"web/src/controllers"
)

var logoutRoutes = Route{
	URI:                "/logout",
	Method:             http.MethodGet,
	Handler:            controllers.LogoutUser,
	NeedAuthentication: true,
}
