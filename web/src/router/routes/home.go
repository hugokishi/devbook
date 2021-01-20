package routes

import (
	"net/http"
	"web/src/controllers"
)

var homeRoutes = Route{
	URI:            "/home",
	Method:         http.MethodGet,
	Handler:        controllers.LoadHomePage,
	Authentication: true,
}
