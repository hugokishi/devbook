package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON - Generic Response in JSON
func JSON(w http.ResponseWriter, statusCode int, datas interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if datas != nil {
		if err := json.NewEncoder(w).Encode(datas); err != nil {
			log.Fatal(err)
		}
	}
}

// Error - Return error in JSON format
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
