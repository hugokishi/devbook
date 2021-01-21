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
}
