package controllers

import (
	"net/http"
	"web/src/cookies"
)

// LogoutUser - Logout user in web
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
