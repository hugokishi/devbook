package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/cookies"
	"web/src/requests"
	"web/src/responses"

	"github.com/gorilla/mux"
)

// CreateUser - Call api to create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UnfollowUser - Unfollow user in API
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// FollowUser - Follow user in api
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdateProfile - call api to edit user
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePassword - Update user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, err := json.Marshal(map[string]string{
		"now": r.FormValue("now"),
		"new": r.FormValue("new"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/change-password", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(passwords))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// DeleteUser - Delete user in API
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, err := requests.RequestsWithAuthentication(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
