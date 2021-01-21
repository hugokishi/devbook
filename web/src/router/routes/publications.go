package routes

import (
	"net/http"
	"web/src/controllers"
)

var publicationsRoutes = []Route{
	{
		URI:                "/publications",
		Method:             http.MethodPost,
		Handler:            controllers.CreatePublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}/like",
		Method:             http.MethodPost,
		Handler:            controllers.LikePublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}/deslike",
		Method:             http.MethodPost,
		Handler:            controllers.DeslikePublication,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}/edit",
		Method:             http.MethodGet,
		Handler:            controllers.LoadEditPage,
		NeedAuthentication: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Method:             http.MethodPut,
		Handler:            controllers.UpdatePublication,
		NeedAuthentication: true,
	},
}
