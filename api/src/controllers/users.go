package controllers

import (
	"net/http"
)

// CreateUser - Function for create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

// ListUsers - Function for list all users in databbase
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all Users"))
}

// ListUser - Function for list one user in database
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List one User"))
}

// UpdateUser - Function for update the user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

// DeleteUser - Function for delete the user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
