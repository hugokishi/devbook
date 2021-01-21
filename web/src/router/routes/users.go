package routes

import (
	"net/http"
	"web/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                "/register",
		Method:             http.MethodGet,
		Handler:            controllers.LoadRegisterPage,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Handler:            controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/search-users",
		Method:             http.MethodGet,
		Handler:            controllers.LoadUsersPage,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Handler:            controllers.LoadUserPage,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userId}/unfollow",
		Method:             http.MethodPost,
		Handler:            controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userId}/unfollow",
		Method:             http.MethodPost,
		Handler:            controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userId}/follow",
		Method:             http.MethodPost,
		Handler:            controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/profile",
		Method:             http.MethodGet,
		Handler:            controllers.LoadUserLoggedProfile,
		NeedAuthentication: true,
	},
	{
		URI:                "/edit-profile",
		Method:             http.MethodGet,
		Handler:            controllers.LoadUserEditPage,
		NeedAuthentication: true,
	},
	{
		URI:                "/edit-profile",
		Method:             http.MethodPut,
		Handler:            controllers.UpdateProfile,
		NeedAuthentication: true,
	},
	{
		URI:                "/edit-password",
		Method:             http.MethodGet,
		Handler:            controllers.LoadUserPasswordEditPage,
		NeedAuthentication: true,
	},
	{
		URI:                "/edit-password",
		Method:             http.MethodPost,
		Handler:            controllers.UpdatePassword,
		NeedAuthentication: true,
	},
	{
		URI:                "/delete-user",
		Method:             http.MethodDelete,
		Handler:            controllers.DeleteUser,
		NeedAuthentication: true,
	},
}
