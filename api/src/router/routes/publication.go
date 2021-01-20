package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Handler:        controllers.CreatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications",
		Method:         http.MethodGet,
		Handler:        controllers.GetPublications,
		Authentication: true,
	},
	{
		URI:            "/publications/{id}",
		Method:         http.MethodGet,
		Handler:        controllers.GetPublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{id}",
		Method:         http.MethodPut,
		Handler:        controllers.UpdatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{id}",
		Method:         http.MethodDelete,
		Handler:        controllers.DeletePublication,
		Authentication: true,
	},
}
