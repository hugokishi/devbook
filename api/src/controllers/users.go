package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ListUsers - Function for list all users in databbase
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all Users"))
}

// CreateUser - Function for create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(request, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)

	userID, err := repository.CreateNewUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID Inserted: %d", userID)))

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
