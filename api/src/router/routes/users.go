package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Handler:        controllers.ListUsers,
		Authentication: true,
	},
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Handler:        controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Handler:        controllers.ListUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Handler:        controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Handler:        controllers.DeleteUser,
		Authentication: true,
	},
}
