package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Login - Method to authenticate the user in API
func Login(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(request, &user); err != nil {
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
	userInDb, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(userInDb.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.GenerateToken(userInDb.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userInDb.ID, 10)

	responses.JSON(w, http.StatusOK, models.DataAuthentication{ID: userID, Token: token})

}
