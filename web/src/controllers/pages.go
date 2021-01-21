package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
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

	utils.RunTemplate(w, "home.html", publications)
}
