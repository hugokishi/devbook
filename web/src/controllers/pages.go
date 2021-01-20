package controllers

import (
	"net/http"
	"web/src/utils"
)

// LoadLoginPage - Render login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "login.html", nil)
}

// LoadRegisterPage - Render register page
func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "register.html", nil)
}
