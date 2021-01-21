package routes

import (
	"net/http"
	"web/src/controllers"
)

var principalRoutes = Route{
	URI:                "/feed",
	Method:             http.MethodGet,
	Handler:            controllers.LoadHomePage,
	NeedAuthentication: true,
}
