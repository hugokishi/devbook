package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
