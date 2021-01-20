package router

import (
	"web/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter - Function for generate new Router
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
