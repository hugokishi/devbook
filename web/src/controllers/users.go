package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:3333/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(bytes.NewBuffer(user))
	fmt.Println(response.Body)
}
