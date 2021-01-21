package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
	"web/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage - Render login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/feed", 302)
		return
	}

	utils.RunTemplate(w, "login.html", nil)
}

// LoadRegisterPage - Render register page
func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "register.html", nil)
}

// LoadHomePage - Render home page
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	var publications []models.Publication
	if err := json.NewDecoder(response.Body).Decode(&publications); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.RunTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// LoadEditPage - Load edit page
func LoadEditPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	var publication models.Publication
	if err = json.NewDecoder(response.Body).Decode(&publication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.RunTemplate(w, "edit-publication.html", publication)
}

// LoadUsersPage - Load users page
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?usuario=%s", config.APIURL, nameOrNick)

	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	var users []models.User
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.RunTemplate(w, "users.html", users)
}

// LoadUserPage - Load user profile
func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userIDLogged, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchCompleteUser(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	if userID == userIDLogged {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	utils.RunTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedID uint64
	}{
		User:         user,
		UserLoggedID: userIDLogged,
	})
}

// LoadUserLoggedProfile - Load the user profile
func LoadUserLoggedProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchCompleteUser(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.RunTemplate(w, "profile.html", user)

}

// LoadUserEditPage - Load page to edit user
func LoadUserEditPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)

	go models.GetUserData(channel, userID, r)

	user := <-channel
	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: "Erro ao buscar o usuário"})
		return
	}

	utils.RunTemplate(w, "edit-profile.html", user)
}

// LoadUserPasswordEditPage - Method to load page to update password
func LoadUserPasswordEditPage(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "update-password.html", nil)
}
