package routes

import (
	"net/http"
	"web/src/middlewares"

	"github.com/gorilla/mux"
)

// Route - Represent interface for routes
type Route struct {
	URI                string
	Method             string
	Handler            func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

// Setup - Setup the routes in router
func Setup(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, principalRoutes)

	for _, route := range routes {
		if route.NeedAuthentication {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
