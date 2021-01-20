package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route - Represent interface for routes
type Route struct {
	URI            string
	Method         string
	Handler        func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Setup - Setup the routes in router
func Setup(router *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return router
}
