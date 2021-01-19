package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter - Generate new router for application
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
