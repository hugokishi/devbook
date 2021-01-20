package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// ListUsers - Function for list all users in databbase
func ListUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.GetUsers(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// CreateUser - Function for create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(request, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.CreateNewUser(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
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
