package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requests"
)

// User - User type
type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreatedAt    time.Time     `json:"createdAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

// SearchCompleteUser - Search user
func SearchCompleteUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followerChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationChannel := make(chan []Publication)

	go GetUserData(userChannel, userID, r)
	go GetFollower(followerChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	go GetPublication(publicationChannel, userID, r)

	var (
		user         User
		follower     []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("Erro ao buscar o usuário")
			}
			user = loadedUser
		case loadedFollower := <-followerChannel:
			if loadedFollower == nil {
				return User{}, errors.New("Erro ao buscar os seguidores")
			}
			follower = loadedFollower
		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("Erro ao buscar quem o usuário esta seguindo")
			}
			following = loadedFollowing
		case loadedPublication := <-publicationChannel:
			if loadedPublication == nil {
				return User{}, errors.New("Erro ao buscar as publicações")
			}
			publications = loadedPublication
		}
	}

	user.Followers = follower
	user.Following = following
	user.Publications = publications

	return user, nil
}

// GetUserData - Get User Data
func GetUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// GetFollower - Get follower Data
func GetFollower(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// GetFollowing - Get following Data
func GetFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

// GetPublication - Get publication Data
func GetPublication(channel chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)
	response, err := requests.RequestsWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publication
	if err = json.NewDecoder(response.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publication, 0)
		return
	}

	channel <- publications
}
