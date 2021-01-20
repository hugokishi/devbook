package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route - Struct for routes
type Route struct {
	URI            string
	Method         string
	Handler        func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Setup - Setup all routes
func Setup(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoutes)

	for _, route := range routes {
		if route.Authentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	return r
}
