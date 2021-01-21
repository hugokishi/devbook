package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorAPI - Type Error
type ErrorAPI struct {
	Err string `json:"error"`
}

// JSON - Return responses in JSON for requests
func JSON(w http.ResponseWriter, statusCode int, datas interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(datas); err != nil {
			log.Fatal(err)
		}
	}
}

// TreatStatusCode - Treat return of request in range 400 or more
func TreatStatusCode(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
